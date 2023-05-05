# 下载安装

```bash
python -m pip install -U pyautogui
```

# 参数相关

## 参数的设置

**一、自动防故障功能**

* 使用`.FAILSAFE=false`关闭

```python
pyautogui.FAILSAFE=false
```

**二、停顿功能**

* 使`pyautogui`所有操作停顿`1s`
* 获取参数不受该影响（鼠标的位置）

```python
pyautogui.PAUSE = 1
```

## 参数的获取

**一、获取屏幕分辨率**

* 使用`.size()`获取屏幕分辨率
* 返回的值为`width`和`height`

```python
width,height=pyautogui.size()
```

**二、获取鼠标的位置**

* 使用`.position()`获取鼠标位置

**三、GUI窗口实时获取鼠标参数**

* 使用`.mouseInfo()`可以通过GUI窗口实时获取鼠标实时位置

```python
pyautogui.mouseInfo()
```

# 鼠标事件

## 鼠标移动事件

**一、鼠标移动到某点**

* 使用`.moveTo()`实现鼠标移动到某点
* 需要输入3个参数：x位置(`px`)，y位置(`px`)，持续时间（`s`）

```python
#移动到(100,300)处
pyautogui.moveTo(100,300,4)
```

**二、鼠标位移**

* 使用`.moveRel`实现相对位移
* 需要输入3个参数：x位移(`px`)，y位移(`px`)，持续时间（`s`）
* 右下为正，左上为负

```python
#向右下方向位移
pyautogui.moveRel(300,300,4)
```

## 鼠标点击事件

**一、单击**

* 使用`.click()`实现
* 可以输入指定位置，默认为当前位置
* 可以设置点击次数(`clicks`)、点击间隔(`interval`)和移动到指定位置时间(`duration`)

```python
#在(100.100)处点击
pyautogui.click(100,100)
#当前位置点击
pyautogui.click()
```

**二、双击**

* 使用`.doubleClick()`实现

```python
pyautogui.doubleClick(100,100)
```

**三、滚轮点击**

* 使用`.middleClick()`实现

```python
pyautogui.middleClick(100,100)
```

**四、自定义点击**

* 使用`.click()`,`rightClick()`,`middleClick()`实现
* 可以输入指定位置，默认为当前位置
* 可以设置点击次数(`clicks`)、点击间隔(`interval`)和移动到指定位置时间(`duration`)

```python
#移动到(100,100)处，点击3次左键，间隔0.1s，鼠标移动过渡时间为0.5s
pyautogui.click(100,100, clicks=3,interval=0.1,duration=0.5)
```

## 鼠标按住事件

**一、按住**

* 使用`.mouseDown()`实现

```python
pyautogui.mouseDown()
```

**二、释放**

* 使用`.mouseUp()`实现

```python
pyautogui.mouseUp()
```

## 鼠标拖动事件

**一、拖动到指定位置**

* 使用`.dragTo()`实现
* 参数：1、x坐标 2、y坐标 3、事件花费时间

```python
pyautogui.dragTo(100,300,duration=1)
```

**二、相对拖动**

* 使用`.dragRel()`实现
* 右下为正

```python
pyautogui.dragRel(100,500,duration=4)
```

## 鼠标滚轮事件

**一、滚动滚轮**

* 使用`.sroll()`实现
* 正为向下

```python
pyautogui.scroll(300)
```

**二、鼠标移动后滚动**

* 使用`.sroll()`实现
* 输入3个参数：1、滚动量  2、x坐标  3、y坐标

```python
pyautogui.scroll(-10,1000,700)
```

# 键盘事件

**一、按下事件**

* 使用`.press()`实现按键按下
* 每次执行`.press()`，都会执行一次`keyUp()`和`keyDown()`
* 对字母支持大小写
* 可以按特殊控键

```python
#输入a
pyautogui.press('a')
#输入A
pyautogui.press('A')
```

**二、组合键的快速输入**

* 使用字符串数组实现组合键的顺序按下
* `interval`参数不能设置组合键之间的间隔时间

```python
pyautogui.press(['a','b','c'])
```

**三、组合键有间隔时间输入**

* 使用`typewrite()`实现
* 使用`interval`参数设置按键间的间隔

```python
pyautogui.typewrite(['s','r','f'], interval=1)
```

**四、文本输入**

* 使用`typewrite()`实现
* 使用`interval`参数设置字符出现的间隔时间

```python
pyautogui.typewrite('hello, PyAutoGUI!\n')
```

**五、组合键的按住连续输入**

* 使用`.hotkey`实现

```python
#调出任务管理器
pyautogui.hotkey('ctrl', 'shift', 'esc')
pyautogui.hotkey('ctrl','c')
```

**六、特殊按键列表**

| **键盘字符串**                  | **描述**                         |
| ------------------------------- | -------------------------------- |
| enter(或return 或 \n)           | 回车                             |
| esc                             | ESC键                            |
| shiftleft, shiftright           | 左右SHIFT键                      |
| altleft, altright               | 左右ALT键                        |
| ctrlleft, ctrlright             | 左右CTRL键                       |
| tab (\t)                        | TAB键                            |
| backspace, delete               | BACKSPACE 、DELETE键             |
| pageup, pagedown                | PAGE UP 和 PAGE DOWN键           |
| home, end                       | HOME 和 END键                    |
| up, down, left,right            | 箭头键                           |
| f1, f2, f3…. f12                | F1…….F12键                       |
| volumemute, volumedown,volumeup | 声音变大变小静音（有些键盘没有） |
| pause                           | PAUSE键，暂停键                  |
| capslock                        | CAPS LOCK 键                     |
| numlock                         | NUM LOCK 键                      |
| scrolllock                      | SCROLLLOCK 键                    |
| insert                          | INSERT键                         |
| printscreen                     | PRINT SCREEN键                   |
| winleft, winright               | Win键（windows ）                |
| command                         | command键（Mac OS X ）           |
| option                          | option（Mac OS X）               |

# 屏幕事件

**一、截图**

* 使用`.screenshot()`实现截图，返回`Pillow`对象
* 使用`.save()`保存

```python
im =pyautogui.screenshot()
im.save('文件路径/屏幕截图.png')
```

**二、图片匹配获取位置**

* 使用`.locateCenterOnScreen()`获取屏幕上匹配图片的中心位置
* 需要输入1个参数：图片来源，图片越大匹配越慢
* 使用`region`参数可以限制屏幕搜索范围以加快速度
* `region`参数由 左上角x坐标，左上角y坐标、相对宽度、相对高度组成（右下为正）

```python
# region参数限制查找范围，加快查找速度
loc = pyautogui.locateCenterOnScreen("icon_xx.png", region=(0, 0,sizex/2, sizey/10) ) 
pyautogui.moveTo(*loc, duration=0.5) # 移动鼠标
pyautogui.click(clicks=1)
```

# 弹窗

**一、警告弹窗**

* 使用`.alert()`设置警告弹窗
* `text`设置文本
* `title`设置弹窗标题
* `button`设置选项`(OK为确定)`

```python
pyautogui.alert(text='警告',title='PyAutoGUI消息框',button='OK')
```

**二、选择弹窗**

* 使用`.confirm()`实现
* `text`设置文本
* `title`设置弹窗标题
* `buttons`设置选项
* 点击后返回选项的值

```python
msg=pyautogui.confirm(text='请选择',title='PyAutoGUI消息框',buttons=['1','2','3'])
print(msg)
```

**三、文本输入框**

* 使用`.prompt()`实现
* `text`设置文本
* `title`设置弹窗标题
* `default`设置默认值
* 点击确定后返回输入的值

```python
msg=pyautogui.prompt(text='请输入',title='PyAutoGUI消息框',default='请输入')
print(msg)
```

**四、密码框**

* 使用`.password()`实现
* `text`设置文本
* `title`设置弹窗标题
* `mask`设置代替值

```python
 pyautogui.password(text='输入密码',title='PyAutoGUI消息框',default='',mask='*')
```





