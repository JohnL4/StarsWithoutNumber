Sheets, lists, and buttons stay HTML over the canvas. The HUD uses
`pointer-events: none`, and each control that needs clicks sets
`pointer-events: auto`.

The view that mounts the canvas is [web/src/views](../views/README.md). HUD
components use the Options API, in the same form as that sample.
