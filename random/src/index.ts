import { Hono } from 'hono'
import { cors } from 'hono/cors'

interface KeyedSigns {
  [key: string]: string[],
}

interface AllSigns {
  imageCount: number,
  images: string[],
}

type Bindings = {
  signs: KVNamespace
}




const baseUrl = "https://roadsign.pictures/";

const app = new Hono<{Bindings:Bindings}>()
app.use('*', cors())

app.get('/', async (c) => {
  let url = c.req.header("referer") ?? baseUrl
  let signsRawJson = await c.env.signs.get('all');
  let allSigns =  JSON.parse(signsRawJson ?? "{}") as AllSigns
  let newUrl = getRandom(url, allSigns.images);

  return c.redirect(	newUrl, 302);
})

app.get('/state/:state',  async (c) => {
  let url = c.req.header("referer") ?? baseUrl
  let signsRawJson = await c.env.signs.get('state')
  let allSigns =  JSON.parse(signsRawJson ?? "{}") as KeyedSigns

  let {state} = c.req.param()

  if (!state || !(state in allSigns)) {
    return c.notFound()
  }
  const { idonly } = c.req.query()
  console.log(idonly)

  if (idonly) {
    return c.json({imageId: getRandomJson(allSigns[state])});
  } else {
    let newUrl = getRandom(url, allSigns[state]);

    return c.redirect(	newUrl, 302);
  }
})

app.get('/statesubdivision/:county', async (c) => {
  let url = c.req.header("referer") ?? baseUrl
  let signsRawJson = await c.env.signs.get('county');
  let allSigns =  JSON.parse(signsRawJson ?? "{}") as KeyedSigns

  const { county } = c.req.param()


  if (!county || !(county in allSigns)) {
    return c.notFound()
  }

  const { idonly } = c.req.query()

  if (idonly) {
    return c.json({imageId: getRandomJson(allSigns[county])});
  } else {
    let newUrl = getRandom(url, allSigns[county]);

    return c.redirect(	newUrl, 302);
  }
})

app.get('/place/:place',  async (c) => {
  let url = c.req.header("referer") ?? baseUrl
  let signsRawJson = await c.env.signs.get('place');
  let allSigns =  JSON.parse(signsRawJson ?? "{}") as KeyedSigns

  const { place } = c.req.param()

  if (!place || !(place in allSigns)) {
    return c.notFound();
  }

  const { idonly } = c.req.query()
  if (idonly) {
    return c.json({imageId: getRandomJson(allSigns[place])});
  } else {
    let newUrl = getRandom(url, allSigns[place]);

    return c.redirect(	newUrl, 302);
  }
})

const getRandomJson = (images: string[]) => {
  const imageLength = images.length;
  const randIndex = Math.floor(Math.random() * imageLength) as number;
  return images[randIndex];
}

const getRandom = (baseUrl: string, images: string[]) =>  {
  const imageLength = images.length;
  const randIndex = Math.floor(Math.random() * imageLength) as number;
  const url = new URL(baseUrl);
  const newImage = images[randIndex];

  const urlJustHost = new URL(`${url.protocol}//${url.host}`);
  const appendedUrl = new URL(`sign/${newImage}`, urlJustHost);


  return appendedUrl.toString();
}

export default app
