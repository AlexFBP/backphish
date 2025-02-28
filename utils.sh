# List all NS records for given domain. Returns first one
function dns-ns(){
  if [ -z $1 ]; then
    echo 'Domain must be specified'
    echo 'Usage: dns-ns DOMAIN'
    return -1
  fi
  DOMAIN=$1
  dig +noall +answer ns $DOMAIN
  # ble="$(dig +short ns $DOMAIN | head -n 1)"
  # echo $ble
  # return ble
}

function dns-a() {
  if [ -z $1 ]; then
    echo 'Must be specified at least a domain'
    echo 'Usage: dns-a (SUB)DOMAIN [DOMAIN]'
    return -1
  fi
  DOMAIN=''
  DOMAIN_OR_SUBDOMAIN=$1
  if [ -z $2 ]; then
    DOMAIN=$1
  else
    DOMAIN=$2
  fi

  dns-ns $DOMAIN
  # TODO: Detect a debug flag (i.e., -d) to remove '+noall +answer'
  # Maybe with getopt(s) https://aplawrence.com/Unix/getopts.html
  dig +noall +answer @$(dig +short $(dig +short ns $DOMAIN | head -n 1)) $DOMAIN_OR_SUBDOMAIN
  # dig +noall +answer @$(dig +short $(dns-ns $DOMAIN)) $DOMAIN_OR_SUBDOMAIN
}

function bp-single() {
  if [ -z $1 ]; then
    echo 'Unspecified target'
    return 1
  fi

  target=$1
  aditional_args=$2
  ./backphish -l 0 -q 1 -n ${target} $(echo "$2")
}

function bp-retrieve() {
  if [ "$#" -lt 2 ]; then
    echo 'It is required at least the URL and a single pattern'
    return 1
  fi

  target=""
  count=0
  for arg in "$@"; do
    if [ -z "$target" ]; then
      target=$arg
      result=$(curl -s "$target")
      # echo "$result"
      last_status=$?
      if [ $last_status -ne 0 ]; then
        echo "Error while retrieving URL, code: ${last_status}"
        return 2
      fi
    else
      ((count++))
      echo "pattern [${count}] - \"${arg}\" - results:"
      echo "$result" | grep -oP "${arg}"
    fi
  done
}

function bp-retrieve-tokens() {
  if [ -z $1 ]; then
    echo 'Unspecified target'
    return 1
  fi

  target="$1"
  bp-retrieve "${target}" \
    "'https://discord.com/api/webhooks/\K[0-9]+/[a-zA-Z0-9_-]+" \
    "\"\K[0-9]+:[a-zA-Z0-9_-]+" \
    "\"\K-[0-9]+" \
}

function bp-retrieve-tokens-nq9() {
  if [ -z $1 ]; then
    echo 'Unspecified base'
    return 1
  fi

  base="$1"
  pages=("ini" "otp" "dinamica" "dinamica2")
  for page in "${pages[@]}"; do
    bp-retrieve-tokens "${base}/${page}.html"
  done
}
