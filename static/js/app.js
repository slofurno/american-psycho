"use strict";

var el = document.getElementById("apb");

var app = {};

app.vm = (function() {
    var vm = {}
    vm.init = function() {
        vm.sentences = [];
        vm.offset = m.prop(0);
        var count = 0;

        vm.advance = function() {
          console.log("advance");
          var next = vm.offset() + 1;
          vm.offset(next);
        };
/*
        for (var i = 0;i<1000;i++){
          vm.sentences.push(new app.sentence(i));
        }
*/
        vm.add=function(){
          var next = new app.sentence(count);
          count++;
          vm.sentences.push(next);
          m.redraw();
        };

        setInterval(function(){
          vm.add();
        },500);
    }
    return vm;
}());



app.sentence = function(n){
  var s = null;

  var o = {};

  o.getText = function(){

    console.log("gettext");

    if (s !=null){
      return s;
    }

    m.startComputation();
    var deferred = m.deferred();

    var client = new XMLHttpRequest();
    var url = "/api?n=" + n;

    client.onload = function(e){

      var response = JSON.parse(this.response);

      console.log(response);
      s = response.Lines[0];
      deferred.resolve(s);
      m.endComputation();
    };

    client.open("GET",url);
    client.send();

    return deferred.promise;
  };

  return o;
};

app.book = function(){
  var sentences = [];
  var book = {};

  book.getSentence()


};

app.controller = function() {
    app.vm.init();
};

app.view=function(){

  return m("div", {class: "book"}, [

    app.vm.sentences.map(function(sentence){

      return m("div", {class: "sentence"}, sentence.getText());

    })
  ]);
};

m.mount(el, {controller: app.controller, view: app.view});
