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
        function generateToken(form) {
            var form = document.getElementById(form);

            form.submit();
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
                        <input type="text" id="sub" name="sub" value="localhost+ec@rca.fr">
                        <label for="name">Subject :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="id_entreprise" name="id_entreprise" value="1">
                        <label for="name">Id entreprise :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">fiber_pin</i>
                        <input type="text" id="scopes" name="scopes" value="purchase">
                        <label for="name">Scopes :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">fiber_pin</i>
                        <input type="text" id="roles" name="roles" value="RCA_CLOUD_EXPERT_COMPTABLE E_COLLECTE_BO_CREA E_CREATION_CREA E_QUESTIONNAIRE_CREA">
                        <label for="name">Roles :</label>
                    </div>
                </div>
                <div class="row">
                    <div class="input-field col s12">
                        <i class="material-icons prefix">account_balance</i>
                        <input type="text" id="rcaPartnerId" name="rcaPartnerId" value="meg-test-interne">
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
