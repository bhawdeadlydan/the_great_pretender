#!/bin/bash
set -e -u

# make sure we're always running from the same directory, regardless of what the 
# current working directory was when this script was run
MY_DIR="$( cd "$( dirname "$0" )" && pwd )"
REPOS=(
  'pretend_web_app'
	'pretend_catalog_service'
	'pretend_pricing_service'
	'pretend_deals_service'
	)

function clone_if_needed {
	local repo_name=$1
	local repo_dir="$MY_DIR/$repo_name"
	local repo_url="git@github.com:ThoughtWorks-AELab/$repo_name.git"

	if [ ! -d $repo_dir ]
	then
	  echo "Looks like you don't have the $repo_name repo. Cloning it for you now..."
	  git clone $repo_url $repo_dir
  fi
}

function pull {
	local repo_name=$1
	local repo_dir="$MY_DIR/$repo_name"
	echo "pulling $repo_name..."
	(cd $repo_dir && git pull)
}

function clone_all {
  for repo in "${REPOS[@]}"
  do
    clone_if_needed $repo
  done
}

function pull_all {
  for repo in "${REPOS[@]}"
  do
    pull $repo
  done
}


valid_commands="\n\tclone\n\tpull"
case "${1:-}" in 

'')
  echo -e "Valid commands are:$valid_commands"
  ;;
clone)
	clone_all
  ;;
pull)
	clone_all
  pull_all
  ;;
*)
  echo -e "Unrecognized command. Valid commands are:$valid_commands"
  ;;
esac
