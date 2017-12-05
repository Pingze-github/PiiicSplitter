
import os
from PIL import Image

# 判定方法，一行像素与上一行的平均亮度值差大于一定阈值

CONFIG = {
    'gate_brightness_diff': 50, # 行平均亮度差阈值
    'gate_split_min_height_rate': 0.1, # 切块最小高度和宽度比
    'output_dir': 'e:/testdata/splits', # 切块最小高度和宽度比
}

# RGB转亮度(305911公式)
def brightness(color):
    return color[0] * 0.30 + color[1] * 0.59 + color[2] * 0.11

def main():
    path = os.path.abspath('e:/testdata/testPiiic1.jpg')
    piiic = Image.open(path)
    width = piiic.width
    height = piiic.height
    CONFIG['gate_split_min_height'] = CONFIG['gate_split_min_height_rate'] * width
    print('读取到图片 {} ({}x{})'.format(path, width, height))
    pix = piiic.load()
    splitYList = [0]
    print('正在查到分割位置...')
    for y in range(1, height):
        diff = 0
        for x in range(width):
            diff += abs(brightness(pix[x,y]) - brightness(pix[x,y - 1]))
        diffAvg = diff / width
        if (diffAvg > CONFIG['gate_brightness_diff']):
            print('查找到分割位置 {}，亮度差值 {}'.format(y, diffAvg))
            splitYList.append(y)
    splitYList.append(height)
    index = 1
    for i in range(1, len(splitYList)):
        if (splitYList[i] - splitYList[i-1] > CONFIG['gate_split_min_height']):
            splitBox = (0, splitYList[i-1], width, splitYList[i])
            split = piiic.crop(splitBox)
            savePath = CONFIG['output_dir'] + '/{}.jpg'.format(index)
            split.save(savePath)
            print('保存 分块{} 到 {}'.format(index, savePath))
            index += 1

main()