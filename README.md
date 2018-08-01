# Public.Bio

Create a simple public bio for yourself. An open source alternative to about.me, Linktree, Linkkle, etc.

![a preview of public.bio](https://i.snap.as/A0hxIru.png)

## Features

* Single-user mode
* Run dynamically or generate a static site

## Getting Started

Right now it's only made for a single user. Edit `sample.json` and then run:

```
go install ./cmd/publicbio
publicbio -u sample.json
```

You'll see your site at `localhost:8080`. Provide a different port with the `-p` option.

### Generate a static site

Use `publicbio` as a static site generator instead of server application with the `-s` flag. HTML goes to stdout, so direct it to the file you want.

```
publicbio -u sample.json -s > bio.html
```

## Development

After updating styles, run `make`.

## Thanks

Thanks to [International](https://cybre.space/@International) for the initial design, and thanks to [Shane](https://mastodon.design/@ZiiX) for the name!
