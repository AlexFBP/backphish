function dns-a() {
  if [ -z $1 ]; then
    echo 'Must be specified at least a domain'
    return -1
  fi
  DOMAIN=$1
  DOMAIN_OR_SUBDOMAIN=''
  if [ -z $2 ]; then
    DOMAIN_OR_SUBDOMAIN=$1
  else
    DOMAIN_OR_SUBDOMAIN=$2
  fi

  # TODO: Detect a debug flag (i.e., -d) to remove '+noall +answer'
  # Maybe with getopt(s) https://aplawrence.com/Unix/getopts.html
  dig +noall +answer @$(dig +short $(dig +short ns $DOMAIN | head -n 1)) $DOMAIN_OR_SUBDOMAIN
}
