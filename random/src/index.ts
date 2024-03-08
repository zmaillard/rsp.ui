import { Hono } from "hono";
import { cors } from "hono/cors";
import { z } from "zod";
import { zValidator } from "@hono/zod-validator";

interface KeyedSigns {
  [key: string]: string[];
}

interface AllSigns {
  imageCount: number;
  images: string[];
}

type Bindings = {
  signs: KVNamespace;
};

const querySchema = z.object({
  idonly: z.string().optional().pipe(z.coerce.boolean().default(false)),
});

const baseUrl = "https://roadsign.pictures/";

const app = new Hono<{ Bindings: Bindings }>();
app.use("*", cors());

app.get("/", async (c) => {
  let url = c.req.header("referer") ?? baseUrl;
  let signsRawJson = await c.env.signs.get("all");
  let allSigns = JSON.parse(signsRawJson ?? "{}") as AllSigns;
  let newUrl = getRandom(url, allSigns.images);

  return c.redirect(newUrl, 302);
});

app.get(
  "/state/:state",
  zValidator("query", querySchema),
  zValidator("param", z.object({ state: z.string() })),
  async (c) => {
    let url = c.req.header("referer") ?? baseUrl;
    let signsRawJson = await c.env.signs.get("state");
    let allSigns = JSON.parse(signsRawJson ?? "{}") as KeyedSigns;

    let { state } = c.req.valid("param");

    if (!state || !(state in allSigns)) {
      return c.notFound();
    }
    const { idonly } = c.req.valid("query");

    if (idonly) {
      return c.json({ imageId: getRandomJson(allSigns[state]) });
    } else {
      let newUrl = getRandom(url, allSigns[state]);

      return c.redirect(newUrl, 302);
    }
  },
);

app.get(
  "/statesubdivision/:county",
  zValidator("query", querySchema),
  zValidator("param", z.object({ county: z.string() })),

  async (c) => {
    let url = c.req.header("referer") ?? baseUrl;
    let signsRawJson = await c.env.signs.get("county");
    let allSigns = JSON.parse(signsRawJson ?? "{}") as KeyedSigns;

    const { county } = c.req.valid("param");

    if (!county || !(county in allSigns)) {
      return c.notFound();
    }

    const { idonly } = c.req.valid("query");

    if (idonly) {
      return c.json({ imageId: getRandomJson(allSigns[county]) });
    } else {
      let newUrl = getRandom(url, allSigns[county]);

      return c.redirect(newUrl, 302);
    }
  },
);

app.get(
  "/place/:place",
  zValidator("query", querySchema),
  zValidator("param", z.object({ place: z.string() })),
  async (c) => {
    let url = c.req.header("referer") ?? baseUrl;
    let signsRawJson = await c.env.signs.get("place");
    let allSigns = JSON.parse(signsRawJson ?? "{}") as KeyedSigns;

    const { place } = c.req.valid("param");

    if (!place || !(place in allSigns)) {
      return c.notFound();
    }

    const { idonly } = c.req.valid("query");
    if (idonly) {
      return c.json({ imageId: getRandomJson(allSigns[place]) });
    } else {
      let newUrl = getRandom(url, allSigns[place]);

      return c.redirect(newUrl, 302);
    }
  },
);

const getRandomJson = (images: string[]) => {
  const imageLength = images.length;
  const randIndex = Math.floor(Math.random() * imageLength) as number;
  return images[randIndex];
};

const getRandom = (baseUrl: string, images: string[]) => {
  const imageLength = images.length;
  const randIndex = Math.floor(Math.random() * imageLength) as number;
  const url = new URL(baseUrl);
  const newImage = images[randIndex];

  const urlJustHost = new URL(`${url.protocol}//${url.host}`);
  const appendedUrl = new URL(`sign/${newImage}`, urlJustHost);

  return appendedUrl.toString();
};

export default app;
