# Public.Bio

Create a simple public bio for yourself.

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
