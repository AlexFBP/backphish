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
