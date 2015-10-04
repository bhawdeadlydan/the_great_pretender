The Great Pretender

a meta-repo that makes it easier to work on the pretend-* ecommerce thingy:

### Step 1
clone this repo

### Step 2
`./go run`

### Step 3
*there is no step 3!!!!!*


## What's this thing do?
It provides some basic scripting to
- create git clones for the various service repos involved in this system, and make sure each service is prepped to run (e.g. do a bundle install for that service's dependencies) (`./go prep`)
- stand up and run the full fleet of services locally with a single command (`./go run`)
- update each service to the latest (`./go pull`)

## What this ./go thing? Is that a golang thing?
No, just a naming collision. All this is written in bash (so far). See these [two](https://www.thoughtworks.com/insights/blog/praise-go-script-part-i) [articles](https://www.thoughtworks.com/insights/blog/praise-go-script-part-ii) for some background on the whole ./go script idea.

## How does this script know how to prep each service?
It doesn't. It just looks for a `tooling/devprep` script inside each repo after cloning the repo and runs it if it can find it. Service maintainers are expected to use that script to automate whatever a dev needs to do after cloning the repo to get to the state that they can run the service (e.g. bundle install, create and migrate dev databases).
