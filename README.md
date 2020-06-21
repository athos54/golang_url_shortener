# Instructions:

## Clone repo

git clone https://github.com/athos54/golang_url_shortener.git

## Compile binary (optional)

`You must have go installed on your machine, if not, you can skip this step and use url_shortener_api on dist folder`

To compile:

```bash
cd src
```

```bash
bash build
```

## Config env variables

```bash
cp .env_example .env
```

Change the variables of .env file

## Run the containers

```bash
docker-compose up -d
```

## Add url to system

**Note:** short_path must be start with slash like `/my_short_word`

```bash
curl -vvv -X POST --header 'content-type: application/json' -d '{"short_path":"/1217", "long_url": "https://google.es"}' http://localhost:1323
```

## Check the redirect

You can check the redirect with browser or with curl

```bash
curl -v -L http://localhost:1323/1217 2>&1 | egrep "^> (Host:|GET)"
```