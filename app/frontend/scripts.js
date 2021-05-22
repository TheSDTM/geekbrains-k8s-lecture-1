var getJSON = function(url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'json';
    xhr.onload = function() {
      var status = xhr.status;
      if (status === 200) {
        callback(null, xhr.response);
      } else {
        callback(status, xhr.response);
      }
    };
    xhr.send();
};

getJSON('/api', function(err, data) {
  if (err !== null) {
    alert('Произошла ошибка при получении данных: ' + err);
  } else {
    let hostname = data.hostname;
    let appVersion = data.appVersion;

    let envName = data.envName;
    let secret = data.secret;
    let dataDir = data.dataDir;
    let config = data.config;
    let redisRes = data.redisRes;
    let rabbitRes = data.rabbitRes;
    let postgreRes = data.postgreRes;

    
    document.getElementById("appVersion").innerText = appVersion;
    document.getElementById("hostname").innerText = hostname;

    document.getElementById("envVal").innerText = envName;
    document.getElementById("secret").innerText = secret;
    document.getElementById("configmap").innerText = config;
    document.getElementById("pvc").innerText = dataDir;
    document.getElementById("rabbitmq").innerText = rabbitRes;
    document.getElementById("redis").innerText = redisRes;
    document.getElementById("postgres").innerText = postgreRes;
  }
});