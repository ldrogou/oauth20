package templateoauth

//Resultat page de resultat
var Resultat = `<!DOCTYPE html>
<html lang="en">

<head>
    <title>RCA JWT API</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />

    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>

    <script src="http://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.2/rollups/hmac-sha512.js"></script>
    <script src="http://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.2/components/enc-base64-min.js"></script>
    <script>
        function copy(jwt){
            navigator.clipboard.writeText(jwt)
        }

        function test(jwtID){
            var xhr = new XMLHttpRequest(); 
            xhr.open('GET', 'http://localhost:8080/jwt/refresh/' + jwtID);
            xhr.onreadystatechange = function() {
            if (xhr.readyState === 4) {
                alert(xhr.responseText);
            }
            };
            xhr.send();
        }
    </script>
</head>

<body>
    <div>
        <h1 class="center-align">Composition</h1>
    </div>
    <div class="container">
        <a class="waves-effect waves-light btn tooltipped" data-tooltip="Copy" onclick="copy('{{.JwtProduce }}');" >
                <i class="material-icons center">content_copy</i>
        </a>
        <a class="waves-effect waves-light btn tooltipped" data-tooltip="Refresh" onclick="refresh('{{.JwtID }}');" >
                <i class="material-icons left bottom">refresh</i>Refresh Token
        </a>
        <div class="row">
            <div class="light-blue lighten-5 col s5">
                <span style="width:300px; word-wrap:break-word; display:inline-block;">
                    {{.JwtProduce }}
                </span>
            </div>
            <div class="col s7">

                <ul class="collapsible collapsible-accordion">
                    <li>
                        <div class="collapsible-header"><i class="material-icons">account_box</i>header</div>
                        <div class="collapsible-body" ><pre id="header"></pre></div>
                    </li>
                    <li class="active">
                        <div class="collapsible-header"><i class="material-icons">code</i>payload</div>
                        <div class="collapsible-body"><pre id="payload"></pre></div>
                    </li>
                    <li>
                        <div class="collapsible-header"><i class="material-icons">border_color</i>signature</div>
                        <div class="collapsible-body"><span>{{.Sign}}</span></div>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</body>

<script>
   

    let headerGO = JSON.parse('{{.Header}}')
    let payloadGo = JSON.parse('{{.Payload}}')
    var headerJson = JSON.stringify(headerGO, null, 4)
    var payloadJson = JSON.stringify(payloadGo, null, 4)
    document.getElementById("header").innerHTML = "<pre>" + headerJson + "</pre>"
    document.getElementById("payload").innerHTML = "<pre>" + payloadJson + "</pre>"

    document.addEventListener('DOMContentLoaded', function () {
        var elems = document.querySelectorAll('.collapsible');
        var instances = M.Collapsible.init(elems, {});

        var elemsTt = document.querySelectorAll('.tooltipped');
        var instancesTt = M.Tooltip.init(elemsTt, {});
    });

</script>

</html>`
