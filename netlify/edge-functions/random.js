import file from "./common/images.json" assert { type: "json" };

export default async (request, context) =>{

    const imageLen = file.images.length;
    const rand = getRandomInt(imageLen);

    const item = file.images[rand];

    const newPath = "/sign/" + item;
    return new URL(newPath, request.url);
};

function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}
