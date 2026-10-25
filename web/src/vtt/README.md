This directory is the virtual tabletop. It is the shared table for the
assistant, not a standalone game. It does not import Vue.

`createVtt(host)` mounts one Pixi application on the element the view passes
in and returns an object with `select` and `destroy`. The view calls
`destroy` when it unmounts. The scene reads the campaign snapshot and owns
sprites, the camera, and tweens.
