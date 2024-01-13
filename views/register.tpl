<div class="row">
  <div class="col-md-6">
    <div class="panel panel-default">
      <div class="panel-heading">Register</div>
      <div class="panel-body">
       {{template "flash_error" .}}
        <form action="/register" method="post">
          <div class="form-group">
            <label for="username">Username</label>
            <input type="text" id="username" name="username" class="form-control" placeholder="Username">
          </div>
          <div class="form-group">
            <label for="password">Password</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="Password">
          </div>
          <input type="submit" class="btn btn-sm btn-default" value="Register"> Exist user? Go to <a href="/login">Login</a>
        </form>
      </div>
    </div>
  </div>
</div>