## TCP сервер для чтения логов находится на 3333 порту по умолчанию
# Клиента пока нет, пока только telnet
    Все хранится в ОЗУ, потому после перезагрузки данные стираются


1. #log показать список файлов и информации о файлах
2. +log /path/ ...  добавить новый путь например /home/user/ добавить все файлы в папке user
3. #read /path/file.ext ... чтение файла
4. #view_clients - показать все подключение с информацией о них
5. #my_info - показать вашу информацию
6. #exit - выход / разорвать соединение


# В планах

1. Использовать либой файлы как базу либо sqlite
2. Авторизацию 
3. Поддержку чтения логов из докеров