{{ define "deployInfoTable" }}
<html>
{{ template "header" }}
<body>
<div class="container">
    <div class="row">
        <div class="col-md-12">
            <table class="table table-striped table-condensed">
                <thead>
                <th>Deploy ID</th>
                <th>Deploy Time</th>
                <th>Application</th>
                <th>Environment</th>
                <th>Build</th>
                <th>Git SHA</th>
                </thead>

                <tbody>
                {{range .}}
                <tr>
                    <td>{{.UUID}}</td>
                    <td>{{.DeployTime}}</td>
                    <td>{{.Application}}</td>
                    <td>{{.Environment}}</td>
                    <td>{{.Build}}</td>
                    <td>{{.SHA}}</td>
                </tr>
                {{end}}
                </tbody>
            </table>
        </div>
    </div>
</div>
</body>
</html>
{{end}}