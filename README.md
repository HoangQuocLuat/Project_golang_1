# Project_golang_1

Folder routes gồm 2 folder auth và user
+ Folder auth có routes group chứa routes login và register
+ Folder user chứa routes lấy thông tin người dùng

Folder models gồm các đối tượng được ánh xạ từ bảng trong database

Folder types gồm các đối tượng được nhập từ phía người dùng và hiển thị cho người dùng
+ CustomRes trả về kết quả của Err và User
+ Err dùng để tùy chỉnh những thông tin về lỗi 
+ User trả về thông tin người dùng
