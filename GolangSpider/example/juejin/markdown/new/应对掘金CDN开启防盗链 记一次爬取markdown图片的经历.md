# 应对掘金CDN开启防盗链 记一次爬取markdown图片的经历 #

## 使用markdown写文章有什么好处? ##

* markdown是一种纯文本格式(后缀 `.md` ), 写法简单, 不用考虑排版, 输出的文章样式简洁优雅
* markdown自带开源属性, 一次书写后, 即可在任意支持markdown格式的平台发布 (国内支持的平台有, ` 掘金` , ` 知乎(以文档方式导入)` , ` 简书(原本是最好用的, 最近在走下坡路)` )
* 著名代码托管平台github, 每个代码仓库的说明书 ` README.md` 就是典型的markdown格式

原来我喜欢在 掘金或简书后台 写markdown文章, 然后复制粘贴到 gitbook(前提是gitbook已经和github做了关联), 就可以发布到github仓库, 由于内容很吸引人, 在github收获一波stars(stars相当于点赞)

> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a6521cefbb8b?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

但最近掘金和简书等平台突然宣布, 在自己网站存储的图片不再支持外链, 也就是在其它网站请求本站服务器存储的图片一律404 ! 简书是直接封了外链; 掘金发了一个公告, 延期一周执行;

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a6522150b0ea?imageView2/0/w/1280/h/960/ignore-error/1)

## 怎么办? ##

我只好将md文档保存到本地, 然后根据md保存的源图片信息,使用爬虫爬取图片到本地, 然后将图片上传到github仓库(github仓库支持图片上传, 而且不封外链), 将原图片信息替换为github仓库保存的图片信息

## 首先在github新建一个名为 **[GraphBed]( https://link.juejin.im?target=https%3A%2F%2Fgithub.com%2Fzhaoolee%2FGraphBed )** 的仓库, 用来存储图片 ##

> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a652231f8ebb?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

* 将仓库clone到本地 的 ` /Users/lijianzhao/github` 文件夹

` cd /Users/lijianzhao/github git clone https://github.com/zhaoolee/GraphBed.git 复制代码`
> 
> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a65223290eac?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

并保证 在此文件夹下, 有权限push到github, 权限添加方法 [www.jianshu.com/p/716712278…]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F7167122783b5 )

## 将github已有的.md文章对应的仓库下载到本地(以星聚弃疗榜为例) ##

` git clone https://github.com/zhaoolee/StarsAndClown.git 复制代码`
> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a652243d1d70?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

## 编写python脚本 ` md_images_upload.py` ##

> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a65224a16222?imageView2/0/w/1280/h/960/ignore-error/1)
> 此脚本:
> 
> 
> 
> 

* 能搜索当前目录下所有md文件, 将每个md中的图片爬取到本地, 存放到 ` /Users/lijianzhao/github/GraphBed/images` 目录;
* 图片爬取完成后, 自动将 ` /Users/lijianzhao/github/GraphBed/images` 目录下的所有图片, push到Github
* 使用Github中的新图片地址,替换原图片地址
* 大功告成

` import os import imghdr import re import requests import shutil import git import hashlib ## 用户名 user_name = "zhaoolee" ; ## 仓库名 github_repository = "GraphBed" ; ## git仓库在本机的位置 git_repository_folder = "/Users/lijianzhao/github/GraphBed" ## 存放图片的git文件夹路径 git_images_folder = "/Users/lijianzhao/github/GraphBed/images" ## 设置忽略目录 ignore_dir_list=[ ".git" ] # 设置用户代理头 headers = { # 设置用户代理头(为狼披上羊皮) "User-Agent" : "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36" , } # 根据输入的url输入md5命名 def create_name(src_name): src_name = src_name.encode( "utf-8" ) s = hashlib.md5() s.update(src_name) return s.hexdigest() # 获取当前目录下所有md文件 def get_md_files(md_dir): md_files = []; for root, dirs , files in sorted(os.walk(md_dir)): for file in files: # 获取.md结尾的文件 if (file.endswith( ".md" )): file_path = os.path.join(root, file) print (file_path) #忽略排除目录 need_append = 0 for ignore_dir in ignore_dir_list: if (ignore_dir in file_path.split( "/" ) == True): need_append = 1 if (need_append == 0): md_files.append(file_path) return md_files # 获取网络图片 def get_http_image(image_url): image_info = { "image_url" : "" , "new_image_url" : "" } file_uuid_name = create_name(image_url) image_data = requests.get(image_url, headers=headers).content # 创建临时文件 tmp_new_image_path_and_name = os.path.join(git_images_folder, file_uuid_name) with open(tmp_new_image_path_and_name, "wb+" ) as f: f.write(image_data) img_type = imghdr.what(tmp_new_image_path_and_name) if (img_type == None): img_type = "" else : img_type = "." +img_type # 生成新的名字加后缀 new_image_path_and_name = tmp_new_image_path_and_name+img_type # 重命名图片 os.rename(tmp_new_image_path_and_name, new_image_path_and_name) new_image_url = "https://raw.githubusercontent.com/" + user_name + "/" +github_repository+ "/master/" +git_images_folder.split( "/" )[-1]+ "/" +new_image_path_and_name.split( "/" )[-1] image_info = { "image_url" : image_url, "new_image_url" : new_image_url } print (image_info) return image_info # 获取本地图片 def get_local_image(image_url): image_info = { "image_url" : "" , "new_image_url" : "" } try: # 创建文件名 file_uuid_name = uuid.uuid4().hex # 获取图片类型 img_type = image_url.split( "." )[-1] # 新的图片名和文件后缀 image_name = file_uuid_name+ "." +img_type # 新的图片路径和名字 new_image_path_and_name = os.path.join(git_images_folder, image_name); shutil.copy(image_url, new_image_path_and_name) # 生成url new_image_url = "https://raw.githubusercontent.com/" + user_name + "/" +github_repository+ "/master/" +git_images_folder.split( "/" )[-1]+ "/" +new_image_path_and_name.split( "/" )[-1] # 图片信息 image_info = { "image_url" : image_url, "new_image_url" : new_image_url } print (image_info) return image_info except Exception as e: print (e) return image_info # 爬取单个md文件内的图片 def get_images_from_md_file(md_file): md_content = "" image_info_list = [] with open(md_file, "r+" ) as f: md_content = f.read() image_urls = re.findall(r "!\[.*?\]\((.*?)\)" , md_content) for image_url in image_urls: # 处理本地图片 if (image_url.startswith( "http" ) == False): image_info = get_local_image(image_url) image_info_list.append(image_info) # 处理网络图片 else : # 不爬取svg if (image_url.startswith( "https://img.shields.io" ) == False): try: image_info = get_http_image(image_url) image_info_list.append(image_info) except Exception as e: print (image_url, "无法爬取, 跳过!" ) pass for image_info in image_info_list: md_content = md_content.replace(image_info[ "image_url" ], image_info[ "new_image_url" ]) print ( "替换完成后::" , md_content); md_content = md_content with open(md_file, "w+" ) as f: f.write(md_content) def git_push_to_origin(): # 通过git提交到github仓库 repo = git.Repo(git_repository_folder) print ( "初始化成功" , repo) index = repo.index index.add([ "images/" ]) print ( "add成功" ) index.commit( "新增图片1" ) print ( "commit成功" ) # 获取远程仓库 remote = repo.remote() print ( "远程仓库" , remote); remote.push() print ( "push成功" ) def main(): if (os.path.exists(git_images_folder)): pass else : os.mkdir(git_images_folder) # 获取本目录下所有md文件 md_files = get_md_files( "./" ) # 将md文件依次爬取 for md_file in md_files: # 爬取单个md文件内的图片 get_images_from_md_file(md_file) git_push_to_origin() if __name__ == "__main__" : main() 复制代码`
> 
> 
> 
> 
> ###### 几个优化点: ######
> 
> 
> 
> * 支持md引用本地目录图片的爬取(以后就可以在本地编写markdown文件了, 编写完成后, 运行上述脚本,
> 即可自动将md引用的本地图片上传到github, 同时本地图片的引用地址被github在线图片地址所取代)
> * 为防止图片重名, 使用uuid重命名图片名称(后面发现使用uuid会导致相同的网络图片反复爬取保存,
> 所以后面使用网络图片的url地址对应的md5码为新名称, 即可防止生成内容相同, 名称不同的图片)
> * 爬取本地图片,依然使用uuid重名防止重复(个人命名可能会反复使用 ` 001.png` , ` 002.png` 等常用名称)
> * 对爬取的图片, 进行了类型判断, 自动补充图片扩展名
> 
> 
> 

## 使用方法 ##

* 安装python3

安装方法见 [Python数据挖掘 环境搭建]( https://link.juejin.im?target=https%3A%2F%2Fwww.jianshu.com%2Fp%2F40bb0d9f670c )

* 将脚本 ` md_images_upload.py` 放到 ` /Users/lijianzhao/github/GraphBed` 目录 (这里目录可以按照自己的来, 但脚本顶部的几行参数也要修改)

> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a652455ef041?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a6524a4aa68b?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

* 在命令行安装相关依赖包
` pip3 install requests pip3 install git 复制代码` * 从命令行进入 ` /Users/lijianzhao/github/GraphBed`
` cd /Users/lijianzhao/github/GraphBed 复制代码` * 运行脚本
` python3 md_images_upload.py 复制代码`
> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a6525a3873ae?imageslim) 这里我已经是第二次替换图片了,
> 所以上面的动图显示的原图片也是GitHub的图片, 说明脚本第一次已完全替换成功~
> 
> 
> 
> 

图片又可以显示了

> 
> 
> 
> 
> 
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a65278cd1854?imageView2/0/w/1280/h/960/ignore-error/1)
> ![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a65278b4c93d?imageView2/0/w/1280/h/960/ignore-error/1)
> 
> 
> 
> 
> 

被替换为github图片替换后md在线展示地址: [zhaoolee.gitbooks.io/starsandclo…]( https://link.juejin.im?target=https%3A%2F%2Fzhaoolee.gitbooks.io%2Fstarsandclown%2Fcontent%2F )

![](https://user-gold-cdn.xitu.io/2019/6/6/16b2a66e423e0f7a?imageView2/0/w/1280/h/960/ignore-error/1)