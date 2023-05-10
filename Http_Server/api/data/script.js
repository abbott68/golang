document.addEventListener('DOMContentLoaded', function() {
    // 在页面加载完成后执行的代码
    console.log('Page loaded.');

    // 使用Ajax发送GET请求到Golang后端的API
    fetch('/')
        .then(response => response.json())
        .then(data => {
            console.log('Response from Golang backend:', data);
        })
        .catch(error => console.error('Error:', error));
});
