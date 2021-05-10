
set -euo pipefail

readonly KEEPALIVE_INTERVAL=300 

main() {
  keepalive
  $@
}

keepalive() {
  local child_pid

  repeat "keepalive" &
  child_pid=$!
  ensureChildOnEXIT "${child_pid}"
}

repeat() {
  local this="$1"
  while true; do
    echo "${this}"
    sleep "${KEEPALIVE_INTERVAL}"
  done
}

ensureChildOnEXIT() {

  local child_pid="$1"
  trap "kill ${child_pid}" EXIT
}

main "$@"
