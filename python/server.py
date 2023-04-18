import io
from flask import Flask, send_from_directory
from flask import send_file, request
import webuiapi
from PIL import Image, ImageDraw





app = Flask(__name__)


@app.route('/images/<path:path>')
def serve_image(path):
    return send_from_directory('images', path)


@app.route('/txt2img', methods=['POST'])
def txt2img():
    api = webuiapi.WebUIApi(host='10.11.40.31', port=7860)
    width = request.form.get('width', 512)
    height = request.form.get('height', 512)

    # Check if width or height is 0, then set default value to 512
    if not width.isdigit() or int(width) == 0:
        width = 512
    if not height.isdigit() or int(height) == 0:
        height = 512

    prompt = request.form['prompt']
    negative_prompt = request.form['negative_prompt']
    result1 = api.txt2img(prompt=f"{prompt}",
                          negative_prompt=f"{negative_prompt}",
                          seed=1003,
                          styles=["anime"],
                          cfg_scale=7,
                          width=width,
                          height=height,
                          )

    image_stream = io.BytesIO()
    result1.image.save(image_stream, format='JPEG')
    image_stream.seek(0)
    return send_file(
        image_stream,
        mimetype='image/jpeg'
    )


@app.route('/img2img', methods=['POST'])
def img2img():
    api = webuiapi.WebUIApi(host='10.11.40.31', port=7860)
    width = request.form.get('width', 512)
    height = request.form.get('height', 512)

    # Check if width or height is 0, then set default value to 512
    if not width.isdigit() or int(width) == 0:
        width = 512
    if not height.isdigit() or int(height) == 0:
        height = 512

    file = request.files['file']
    image = Image.open(file)
    prompt = request.form['prompt']
    result1 = api.img2img(
        images=[image],  # Pass the image as a list
        prompt=prompt,
        # negative_prompt=negative_prompt,
        seed=1003,
        styles=["anime"],
        cfg_scale=7,
        width=width,
        height=height,
    )

    image_stream = io.BytesIO()
    result1.image.save(image_stream, format='JPEG')
    image_stream.seek(0)
    return send_file(
        image_stream,
        mimetype='image/jpeg'
    )


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
