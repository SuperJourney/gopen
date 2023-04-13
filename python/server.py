import io
from flask import Flask, send_from_directory
from flask import send_file,request
import webuiapi
from PIL import Image, ImageDraw


api = webuiapi.WebUIApi(host='10.11.40.31', port=7860)


app = Flask(__name__)


@app.route('/images/<path:path>')
def serve_image(path):
    return send_from_directory('images', path)


@app.route('/txt2img', methods=['POST'])
def txt2img():
    prompt = request.form['prompt']
    negative_prompt = request.form['negative_prompt']
    result1 = api.txt2img(prompt=f"{prompt}",
                          negative_prompt=f"{negative_prompt}",
                          seed=1003,
                          styles=["anime"],
                          cfg_scale=7,
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
    # 从请求中获取图片文件
    print(request.files)

    file = request.files['file']
    # 保存图片文件到本地
    file.save('./images/received_image.jpg')
    return '图片上传成功！'


if __name__ == '__main__':
    app.run()


