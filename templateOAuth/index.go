package templateoauth

//TemplateIndex index html
var TemplateIndex = `<!DOCTYPE html>
<html>

<head>
    <title>RCA JWT API</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />

    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>

    <script src="http://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.2/rollups/hmac-sha512.js"></script>
    <script src="http://cdnjs.cloudflare.com/ajax/libs/crypto-js/3.1.2/components/enc-base64-min.js"></script>

    <script>

        function base64url(source) {
            // Encode in classical base64
            encodedSource = CryptoJS.enc.Base64.stringify(source);

            // Remove padding equal characters
            encodedSource = encodedSource.replace(/=+$/, '');

            // Replace characters according to base64url specifications
            encodedSource = encodedSource.replace(/\+/g, '-');
            encodedSource = encodedSource.replace(/\//g, '_');

            return encodedSource;
        }

        function generateToken(form) {
            var form = document.getElementById(form);

            form.submit();

            //window.location = 'https://captation.beta.rca.fr/entreprise-partenaire/authorize?client_id=meg-test-interne&scope=user.read company.read accounting_firm.read sales&current_company=true&redirect_uri=http://localhost:8080/oauth/redirect'
        }

        function generate() {
            var header = {
                "alg": "HS512"
            };

            var data = {
                "sub": document.getElementById('sub').value,
                "exp": Math.floor(Date.now() / 1000) + 6 * 30 * 24 * 3600,
                "roles": [
                    "RCA_CLOUD_EXPERT_COMPTABLE",
                    "E_COLLECTE_BO_CREA",
                    "E_CREATION_CREA",
                    "E_QUESTIONNAIRE_CREA"
                ],
                "id_entreprise": document.getElementById('id_entreprise').value,
                "rcaPartnerId": document.getElementById('rcaPartnerId').value
            };

            var secret = document.getElementById('secret').value;
            secret = CryptoJS.enc.Base64.parse(secret);

            var stringifiedHeader = CryptoJS.enc.Utf8.parse(JSON.stringify(header));
            var encodedHeader = base64url(stringifiedHeader);

            var stringifiedData = CryptoJS.enc.Utf8.parse(JSON.stringify(data));
            var encodedData = base64url(stringifiedData);

            var signature = encodedHeader + "." + encodedData;
            signature = CryptoJS.HmacSHA512(signature, secret);
            signature = base64url(signature);

            document.getElementById('jwt').value = encodedHeader + "." + encodedData + "." + signature;
            M.updateTextFields();
            M.textareaAutoResize(document.getElementById('jwt'));
        }
    </script>
</head>

<body>
    <div>
        <h1 class="center-align">JWT</h1>
      </div>
    <div class="container">
        <div class="row">
            <form class="col s6 light-blue lighten-5" id="formLocal" method="post" action="/local">
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_circle</i>
                        <input type="text" id="sub" name="sub" value="mbola.randriamamonjisoa+ec@rca.fr">
                        <label for="name">Subject :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="id_entreprise" name="id_entreprise" value="85422">
                        <label for="name">Id entreprise :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="rcaPartnerId" name="rcaPartnerId" value="agora-expert">
                        <label for="name" >ID partenaire RCA :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">fiber_pin</i>
                        <input type="text" id="secret" name="secret" value="XXXXXXX">
                        <label for="name">Secret :</label>
                    </div>
                </div>
                <div class="row">
                    <a class="waves-effect waves-light btn" onclick="generateToken('formLocal');"><i
                            class="material-icons left">cloud</i>Local</a>
                </div>
            </form>
        
            <form class="col s6 light-green lighten-5" id="formOAtuh20" method="post" action="/oauth20">
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="domain" name="domain" value="captation.beta.rca.fr">
                        <label for="name">Domaine :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="clientId" name="clientId" value="meg-test-interne">
                        <label for="name">Client Id :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="clientSecret" name="clientSecret" value="xxxxxxxx">
                        <label for="name">Client Secret :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="scopes" name="scopes" value="user">
                        <label for="name">Scopes</label>
                    </div>
                </div>
                <div class="row">
                    <div class="checkbox col s12">
                        <label>
                            <input type="checkbox" id="currentCompany" name="currentCompany" checked="checked" />
                            <span>Company courante</span>
                        </label>
                    </div>
                </div>
                <div class="row">
                    <a class="waves-effect waves-light btn" onclick="generateToken('formOAtuh20');"><i
                            class="material-icons left">cloud</i>OAuth2.0</a>
                </div>
            </form>
        </div>

    </div>

</body>

</html>`