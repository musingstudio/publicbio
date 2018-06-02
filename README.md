# Public.Bio

Create a simple public bio for yourself.

![a preview of public.bio](https://i.snap.as/A0hxIru.png)

Right now it's only made for a single user. Edit `sample.json` and then run:

```
cd cmd/publicbio
go build
./publicbio -u ../../sample.json
```

You'll see your site at `localhost:8080`. Provide a different port with the `-p` option.

## Development

After updating styles, run `make`.
