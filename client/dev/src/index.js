if(process.env.NODE_ENV === "production"){
	__webpack_public_path__ = "/static/"
}else{
	__webpack_public_path__ = "http://localhost:3000/static/"
}

require.ensure([], require => {
  require('utils/global');
  ReactDOM.render(require('./router').default, document.getElementById('app'));
}, 'client');
