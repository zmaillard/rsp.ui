export default async (request, context) =>{
    let tempImages = {
        "images": [
            "957569704",
            "957572164",
            "957574672",
            "957579738",
            "957584716",
            "957589126",
            "957591378",
            "957593270",
            "957599628",
            "957617904",
            "957620242",
            "957622472",
            "957624534"
        ]
    }

    const imageLen = tempImages.images.length;
    const rand = getRandomInt(imageLen);

    const item = tempImages.images[rand];

    const url = new URL("/sign/" + item);
    return Response.redirect(url);
};

function getRandomInt(max) {
    return Math.floor(Math.random() * max);
}

export const config = {
    path: "/random",
};