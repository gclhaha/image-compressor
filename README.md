# 图片压缩和尺寸调整工具

这个项目旨在提供一个能够批量处理图片文件的工具，主要包括图片压缩和尺寸调整功能。你可以指定一个输入目录，程序会遍历该目录下的所有图片文件（支持jpg/jpeg和png格式），将它们压缩到指定的文件大小（默认1000KB）并调整宽度不超过指定的最大值（默认1920像素），然后将处理后的图片保存到指定的输出目录中，保持原始目录结构。

## 使用方法

[直接下载](https://github.com/gclhaha/image-compressor/releases)  

或

1. **安装Go编程语言**
   - 确保你的计算机上安装了Go编程语言的环境。

2. **下载和编译程序**
   - 下载项目代码并使用以下命令编译：

     Windows:

     ```bash
     GOOS=windows GOARCH=amd64 go build -o image-compressor.exe
     ```

     Linux:

     ```bash
     GOOS=linux GOARCH=amd64 go build -o image-compressor
     ```

     Mac:

     ```bash
     GOOS=darwin GOARCH=amd64 go build -o image-compressor
     ```

3. **运行程序**
   - 使用命令行参数运行程序，以下是参数的说明：
     - `-i 输入目录路径`：必填，指定包含待处理图片的输入目录路径。
     - `-o 输出目录路径`：必填，指定处理后图片保存的输出目录路径。
     - `-s 目标大小（KB）`：可选，指定压缩后的目标文件大小，默认为1000KB。
     - `-w 最大宽度`：可选，指定图片调整后的最大宽度，默认为1920像素。

   - 示例：
  
     ```bash
     ./image-compressor -i /path/to/input/dir -o /path/to/output/dir -s 1200 -w 1600
     ```

4. **查看处理结果**
   - 程序会将处理后的图片保存在指定的输出目录中，保持原始目录结构。处理过程中会输出每张图片的处理状态信息。

5. **注意事项**
   - 确保输入目录中包含需要处理的图片文件。
   - 程序目前支持jpg/jpeg和png格式的图片处理。
  
