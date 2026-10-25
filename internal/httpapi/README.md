Handlers in this package accept JSON commands, call the domain, and save
through the store. After a commit, the server broadcasts the whole campaign
snapshot to the other connected clients.

The Pinia store that receives that broadcast is
[web/src/stores](../../web/src/stores/README.md).
