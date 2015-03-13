    $(function () {
        var Item = Backbone.Model.extend({
            defaults:function(){
                return {
                  title:"default",
                  url:"default"
                };
            }
        });
        var ItemList = Backbone.Collection.extend({
            model:Item
        });
        var Items = new ItemList;

        var ItemView = Backbone.View.extend({
          tagName: "li",
          template:_.template($('#item-template').html()),
          initialize:function(){
          },
          render:function(){
            this.$el.html(this.template(this.model.toJSON()));
            return this;
          },
          clear:function(){
            this.model.destroy();
          }

        });
        var AppView = Backbone.View.extend({
          el:$("#zlistapp"),
          events:{
            "click .item-source": "showItemList"
          },
          initialize:function(){
            // alert("hha");
            this.listenTo(Items, 'add', this.addOne);
            this.listenTo(Items, 'reset', this.addAll);
            this.listenTo(Items, 'all', this.render);            
            Items.url='/jianshu/now';
            Items.fetch({reset:true});
          },
          addOne:function(item){
            var view = new ItemView({model:item});
            this.$("#item-list").append(view.render().el);
          },
          addAll:function(){
            this.$("#item-list").empty();
            Items.each(this.addOne, this);
          },
          render:function(){
            // alert("render");
          },
          showItemList:function(ev){
            Items.url=$(ev.currentTarget).data('link');
            Items.fetch({reset:true});
            $(".am-dropdown").dropdown("close");
          }

        });
        var App = new AppView;
    });