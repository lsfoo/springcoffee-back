{{define "printq"}}
<!DOCTYPE html>
<html lang="zh">
<head>
<meta charset="UTF-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no" />
<link href="https://cdn.bootcss.com/bootstrap/4.0.0-alpha.6/css/bootstrap.min.css" rel="stylesheet"/>
<title>待打印列表</title>
</head>
<body>
<div class="container">
<div class="starter-template">
<h1>打印</h1>
{{range .}}
<div class="card">
<img class="card-img-top" data-src="http://lfo.oss-cn-beijing.aliyuncs.com/{{.Src}}/" alt="Card image cap">
<div class="card-block">
<h4 class="card-title">{{.NickName}}</h4>
<p class="card-text">{{.Words}}</p>
<p class="card-text">{{.CreateTime}}</p>
<button class="btn btn-outline-primary">打印</button>
<button class="btn btn-outline-secondary">取消</button>
</div>
</div>
{{end}}
</div>
</div>
<script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
<script src="https://cdn.bootcss.com/bootstrap/4.0.0-alpha.6/js/bootstrap.min.js"></script>
</body>
</html>
{{end}}

