<div class="row">
  <div class="col-md-6">
    <div class="panel panel-default">
      <div class="panel-heading">Login</div>
      <div class="panel-body">
        {{template "flash_error" .}}
        <form action="/login" method="post">
          <div class="form-group">
            <label for="username">Username</label>
            <input type="text" id="username" name="username" class="form-control" placeholder="Username">
          </div>
          <div class="form-group">
            <label for="password">Password</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="Password">
          </div>
          <input type="submit" class="btn btn-default" value="Login"> New user? Go to <a href="/register">Register</a>
        </form>
      </div>
    </div>
  </div>
</div>