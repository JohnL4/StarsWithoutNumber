Rules live here. This package does not import HTTP or SQL.

`Roll` returns a `Character` with each attribute rolled on 3d6. `dice` stays
unexported, so callers do not roll individual attributes. The struct's JSON
tags match the field names `ST`, `DX`, `CO`, `IQ`, `WI`, and `CH`.

The matching TypeScript shape is
[web/src/api/character.ts](../../web/src/api/character.ts).
