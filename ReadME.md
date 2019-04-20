## GQLGEN GraphQL Demo using Twitter API [In Progress ... ]

This is a demo for graphql demo using [GQLGEN](http://gqlgen.com). I wanted to use a database for this demo, but then i thought, twitter would be better for such a task.  

### Prerequisites

You will need to have a Twitter Developers App, which you can learn more about [here](https://developer.twitter.com/en/docs/basics/getting-started). Once you have created an App, get the Keys and Tokens for your app. 

Then, clone this repository and then at the root of this repo, add a `.env` file with the following content.

```env
CONSUMER_KEY=<KEY>
CONSUMER_SECRET=<KEY>
ACCESS_TOKEN=<KEY>
ACCESS_SECRET=<KEY>
```

**NB:** Replace the `<KEY>` with the corresponding key from your Twitter App, inside the `.env` file you just created.

### How to run this demo

The easiest way to run this demo is to use docker-compose, so make sure you have [docker](https://docs.docker.com/install/) installed on your device. And then run `docker-conpose up -d` to build and run the graphql server.

After that, open GraphQL Playground on your browser: http://localhost:8080

### TODO

- [ ] Add Dataloader
- [ ] Add Authentication (Not sure how yet)
- [ ] Documentation
