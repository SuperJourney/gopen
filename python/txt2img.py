import webuiapi
import flask

from PIL import Image, ImageDraw

api = webuiapi.WebUIApi(host='10.11.40.31', port=7860)

result1 = api.txt2img(prompt="cute squirrel",
                    negative_prompt="ugly, out of frame",
                    seed=1003,
                    styles=["anime"],
                    cfg_scale=7,
                    )

# draw = ImageDraw.Draw(result1.image)
result1.image.save("e.png")