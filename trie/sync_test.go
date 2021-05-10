
package trie

import (
	"bytes"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
)


func makeTestTrie() (*Database, *SecureTrie, map[string][]byte) {
	// Create an empty trie
	triedb := NewDatabase(memorydb.New())
	trie, _ := NewSecure(common.Hash{}, triedb)

	// Fill it with some arbitrary data
	content := make(map[string][]byte)
	for i := byte(0); i < 255; i++ {
		// Map the same data under multiple keys
		key, val := common.LeftPadBytes([]byte{1, i}, 32), []byte{i}
		content[string(key)] = val
		trie.Update(key, val)

		key, val = common.LeftPadBytes([]byte{2, i}, 32), []byte{i}
		content[string(key)] = val
		trie.Update(key, val)

		// Add some other data to inflate the trie
		for j := byte(3); j < 13; j++ {
			key, val = common.LeftPadBytes([]byte{j, i}, 32), []byte{j, i}
			content[string(key)] = val
			trie.Update(key, val)
		}
	}
	trie.Commit(nil)

	// Return the generated trie
	return triedb, trie, content
}


func checkTrieContents(t *testing.T, db *Database, root []byte, content map[string][]byte) {

	trie, err := NewSecure(common.BytesToHash(root), db)
	if err != nil {
		t.Fatalf("failed to create trie at %x: %v", root, err)
	}
	if err := checkTrieConsistency(db, common.BytesToHash(root)); err != nil {
		t.Fatalf("inconsistent trie at %x: %v", root, err)
	}
	for key, val := range content {
		if have := trie.Get([]byte(key)); !bytes.Equal(have, val) {
			t.Errorf("entry %x: content mismatch: have %x, want %x", key, have, val)
		}
	}
}


func checkTrieConsistency(db *Database, root common.Hash) error {

	trie, err := NewSecure(root, db)
	if err != nil {
		return nil // Consider a non existent state consistent
	}
	it := trie.NodeIterator(nil)
	for it.Next(true) {
	}
	return it.Error()
}


func TestEmptySync(t *testing.T) {
	dbA := NewDatabase(memorydb.New())
	dbB := NewDatabase(memorydb.New())
	emptyA, _ := New(common.Hash{}, dbA)
	emptyB, _ := New(emptyRoot, dbB)

	for i, trie := range []*Trie{emptyA, emptyB} {
		sync := NewSync(trie.Hash(), memorydb.New(), nil, NewSyncBloom(1, memorydb.New()))
		if nodes, paths, codes := sync.Missing(1); len(nodes) != 0 || len(paths) != 0 || len(codes) != 0 {
			t.Errorf("test %d: content requested for empty trie: %v, %v, %v", i, nodes, paths, codes)
		}
	}
}


func TestIterativeSyncIndividual(t *testing.T)       { testIterativeSync(t, 1, false) }
func TestIterativeSyncBatched(t *testing.T)          { testIterativeSync(t, 100, false) }
func TestIterativeSyncIndividualByPath(t *testing.T) { testIterativeSync(t, 1, true) }
func TestIterativeSyncBatchedByPath(t *testing.T)    { testIterativeSync(t, 100, true) }

func testIterativeSync(t *testing.T, count int, bypath bool) {

	srcDb, srcTrie, srcData := makeTestTrie()

	diskdb := memorydb.New()
	triedb := NewDatabase(diskdb)
	sched := NewSync(srcTrie.Hash(), diskdb, nil, NewSyncBloom(1, diskdb))

	nodes, paths, codes := sched.Missing(count)
	var (
		hashQueue []common.Hash
		pathQueue []SyncPath
	)
	if !bypath {
		hashQueue = append(append(hashQueue[:0], nodes...), codes...)
	} else {
		hashQueue = append(hashQueue[:0], codes...)
		pathQueue = append(pathQueue[:0], paths...)
	}
	for len(hashQueue)+len(pathQueue) > 0 {
		results := make([]SyncResult, len(hashQueue)+len(pathQueue))
		for i, hash := range hashQueue {
			data, err := srcDb.Node(hash)
			if err != nil {
				t.Fatalf("failed to retrieve node data for hash %x: %v", hash, err)
			}
			results[i] = SyncResult{hash, data}
		}
		for i, path := range pathQueue {
			data, _, err := srcTrie.TryGetNode(path[0])
			if err != nil {
				t.Fatalf("failed to retrieve node data for path %x: %v", path, err)
			}
			results[len(hashQueue)+i] = SyncResult{crypto.Keccak256Hash(data), data}
		}
		for _, result := range results {
			if err := sched.Process(result); err != nil {
				t.Fatalf("failed to process result %v", err)
			}
		}
		batch := diskdb.NewBatch()
		if err := sched.Commit(batch); err != nil {
			t.Fatalf("failed to commit data: %v", err)
		}
		batch.Write()

		nodes, paths, codes = sched.Missing(count)
		if !bypath {
			hashQueue = append(append(hashQueue[:0], nodes...), codes...)
		} else {
			hashQueue = append(hashQueue[:0], codes...)
			pathQueue = append(pathQueue[:0], paths...)
		}
	}

	checkTrieContents(t, triedb, srcTrie.Hash().Bytes(), srcData)
}

func TestIterativeDelayedSync(t *testing.T) {

	srcDb, srcTrie, srcData := makeTestTrie()

	diskdb := memorydb.New()
	triedb := NewDatabase(diskdb)
	sched := NewSync(srcTrie.Hash(), diskdb, nil, NewSyncBloom(1, diskdb))

	nodes, _, codes := sched.Missing(10000)
	queue := append(append([]common.Hash{}, nodes...), codes...)

	for len(queue) > 0 {

		results := make([]SyncResult, len(queue)/2+1)
		for i, hash := range queue[:len(results)] {
			data, err := srcDb.Node(hash)
			if err != nil {
				t.Fatalf("failed to retrieve node data for %x: %v", hash, err)
			}
			results[i] = SyncResult{hash, data}
		}
		for _, result := range results {
			if err := sched.Process(result); err != nil {
				t.Fatalf("failed to process result %v", err)
			}
		}
		batch := diskdb.NewBatch()
		if err := sched.Commit(batch); err != nil {
			t.Fatalf("failed to commit data: %v", err)
		}
		batch.Write()

		nodes, _, codes = sched.Missing(10000)
		queue = append(append(queue[len(results):], nodes...), codes...)
	}

	checkTrieContents(t, triedb, srcTrie.Hash().Bytes(), srcData)
}

func TestIterativeRandomSyncIndividual(t *testing.T) { testIterativeRandomSync(t, 1) }
func TestIterativeRandomSyncBatched(t *testing.T)    { testIterativeRandomSync(t, 100) }

func testIterativeRandomSync(t *testing.T, count int) {

	srcDb, srcTrie, srcData := makeTestTrie()


	diskdb := memorydb.New()
	triedb := NewDatabase(diskdb)
	sched := NewSync(srcTrie.Hash(), diskdb, nil, NewSyncBloom(1, diskdb))

	queue := make(map[common.Hash]struct{})
	nodes, _, codes := sched.Missing(count)
	for _, hash := range append(nodes, codes...) {
		queue[hash] = struct{}{}
	}
	for len(queue) > 0 {

		results := make([]SyncResult, 0, len(queue))
		for hash := range queue {
			data, err := srcDb.Node(hash)
			if err != nil {
				t.Fatalf("failed to retrieve node data for %x: %v", hash, err)
			}
			results = append(results, SyncResult{hash, data})
		}

		for _, result := range results {
			if err := sched.Process(result); err != nil {
				t.Fatalf("failed to process result %v", err)
			}
		}
		batch := diskdb.NewBatch()
		if err := sched.Commit(batch); err != nil {
			t.Fatalf("failed to commit data: %v", err)
		}
		batch.Write()

		queue = make(map[common.Hash]struct{})
		nodes, _, codes = sched.Missing(count)
		for _, hash := range append(nodes, codes...) {
			queue[hash] = struct{}{}
		}
	}

	checkTrieContents(t, triedb, srcTrie.Hash().Bytes(), srcData)
}


func TestIterativeRandomDelayedSync(t *testing.T) {

	srcDb, srcTrie, srcData := makeTestTrie()

	diskdb := memorydb.New()
	triedb := NewDatabase(diskdb)
	sched := NewSync(srcTrie.Hash(), diskdb, nil, NewSyncBloom(1, diskdb))

	queue := make(map[common.Hash]struct{})
	nodes, _, codes := sched.Missing(10000)
	for _, hash := range append(nodes, codes...) {
		queue[hash] = struct{}{}
	}
	for len(queue) > 0 {

		results := make([]SyncResult, 0, len(queue)/2+1)
		for hash := range queue {
			data, err := srcDb.Node(hash)
			if err != nil {
				t.Fatalf("failed to retrieve node data for %x: %v", hash, err)
			}
			results = append(results, SyncResult{hash, data})

			if len(results) >= cap(results) {
				break
			}
		}

		for _, result := range results {
			if err := sched.Process(result); err != nil {
				t.Fatalf("failed to process result %v", err)
			}
		}
		batch := diskdb.NewBatch()
		if err := sched.Commit(batch); err != nil {
			t.Fatalf("failed to commit data: %v", err)
		}
		batch.Write()
		for _, result := range results {
			delete(queue, result.Hash)
		}
		nodes, _, codes = sched.Missing(10000)
		for _, hash := range append(nodes, codes...) {
			queue[hash] = struct{}{}
		}
	}

	checkTrieContents(t, triedb, srcTrie.Hash().Bytes(), srcData)
}

func TestDuplicateAvoidanceSync(t *testing.T) {

	srcDb, srcTrie, srcData := makeTestTrie()

	diskdb := memorydb.New()
	triedb := NewDatabase(diskdb)
	sched := NewSync(srcTrie.Hash(), diskdb, nil, NewSyncBloom(1, diskdb))

	nodes, _, codes := sched.Missing(0)
	queue := append(append([]common.Hash{}, nodes...), codes...)
	requested := make(map[common.Hash]struct{})

	for len(queue) > 0 {
		results := make([]SyncResult, len(queue))
		for i, hash := range queue {
			data, err := srcDb.Node(hash)
			if err != nil {
				t.Fatalf("failed to retrieve node data for %x: %v", hash, err)
			}
			if _, ok := requested[hash]; ok {
				t.Errorf("hash %x already requested once", hash)
			}
			requested[hash] = struct{}{}

			results[i] = SyncResult{hash, data}
		}
		for _, result := range results {
			if err := sched.Process(result); err != nil {
				t.Fatalf("failed to process result %v", err)
			}
		}
		batch := diskdb.NewBatch()
		if err := sched.Commit(batch); err != nil {
			t.Fatalf("failed to commit data: %v", err)
		}
		batch.Write()

		nodes, _, codes = sched.Missing(0)
		queue = append(append(queue[:0], nodes...), codes...)
	}

	checkTrieContents(t, triedb, srcTrie.Hash().Bytes(), srcData)
}

func TestIncompleteSync(t *testing.T) {

	srcDb, srcTrie, _ := makeTestTrie()


	diskdb := memorydb.New()
	triedb := NewDatabase(diskdb)
	sched := NewSync(srcTrie.Hash(), diskdb, nil, NewSyncBloom(1, diskdb))

	var added []common.Hash

	nodes, _, codes := sched.Missing(1)
	queue := append(append([]common.Hash{}, nodes...), codes...)

	for len(queue) > 0 {

		results := make([]SyncResult, len(queue))
		for i, hash := range queue {
			data, err := srcDb.Node(hash)
			if err != nil {
				t.Fatalf("failed to retrieve node data for %x: %v", hash, err)
			}
			results[i] = SyncResult{hash, data}
		}

		for _, result := range results {
			if err := sched.Process(result); err != nil {
				t.Fatalf("failed to process result %v", err)
			}
		}
		batch := diskdb.NewBatch()
		if err := sched.Commit(batch); err != nil {
			t.Fatalf("failed to commit data: %v", err)
		}
		batch.Write()
		for _, result := range results {
			added = append(added, result.Hash)
		}
		for _, root := range added {
			if err := checkTrieConsistency(triedb, root); err != nil {
				t.Fatalf("trie inconsistent: %v", err)
			}
		}

		nodes, _, codes = sched.Missing(1)
		queue = append(append(queue[:0], nodes...), codes...)
	}

	for _, node := range added[1:] {
		key := node.Bytes()
		value, _ := diskdb.Get(key)

		diskdb.Delete(key)
		if err := checkTrieConsistency(triedb, added[0]); err == nil {
			t.Fatalf("trie inconsistency not caught, missing: %x", key)
		}
		diskdb.Put(key, value)
	}
}

func TestSyncOrdering(t *testing.T) {

	srcDb, srcTrie, srcData := makeTestTrie()

	diskdb := memorydb.New()
	triedb := NewDatabase(diskdb)
	sched := NewSync(srcTrie.Hash(), diskdb, nil, NewSyncBloom(1, diskdb))

	nodes, paths, _ := sched.Missing(1)
	queue := append([]common.Hash{}, nodes...)
	reqs := append([]SyncPath{}, paths...)

	for len(queue) > 0 {
		results := make([]SyncResult, len(queue))
		for i, hash := range queue {
			data, err := srcDb.Node(hash)
			if err != nil {
				t.Fatalf("failed to retrieve node data for %x: %v", hash, err)
			}
			results[i] = SyncResult{hash, data}
		}
		for _, result := range results {
			if err := sched.Process(result); err != nil {
				t.Fatalf("failed to process result %v", err)
			}
		}
		batch := diskdb.NewBatch()
		if err := sched.Commit(batch); err != nil {
			t.Fatalf("failed to commit data: %v", err)
		}
		batch.Write()

		nodes, paths, _ = sched.Missing(1)
		queue = append(queue[:0], nodes...)
		reqs = append(reqs, paths...)
	}

	checkTrieContents(t, triedb, srcTrie.Hash().Bytes(), srcData)

	for i := 0; i < len(reqs)-1; i++ {
		if len(reqs[i]) > 1 || len(reqs[i+1]) > 1 {
			t.Errorf("Invalid request tuples: len(%v) or len(%v) > 1", reqs[i], reqs[i+1])
		}
		if bytes.Compare(compactToHex(reqs[i][0]), compactToHex(reqs[i+1][0])) > 0 {
			t.Errorf("Invalid request order: %v before %v", compactToHex(reqs[i][0]), compactToHex(reqs[i+1][0]))
		}
	}
}
