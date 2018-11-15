<template>
<el-main>
<el-row type="flex" class="row-bg" justify="center">
<el-col :span="8">
</el-col>
<el-col :span="8">
  <el-form ref="form" :model="form" label-width="120px">
  <el-form-item label="Title">
    <el-input placeholder="Insert ToDo title" v-model="title"></el-input>
  </el-form-item>
  <el-form-item label="Description">
    <el-input placeholder="Insert ToDo Description" v-model="description"></el-input>
  </el-form-item>
  <el-form-item>
  <el-button type="primary" @click="addTodo" v-if="selectedTodo!==null">Save</el-button>
  <el-button type="primary" @click="addTodo" v-if="selectedTodo===null">Add</el-button>
  </el-form-item>
  </el-form>
  <ul>
    <li v-for="todo in todoList"><b @click="showTodo(todo)">{{todo.title}}</b> <small>{{todo.description}}</small>
    <el-button @click="deleteTodo(todo)" class="btn">X</el-button>
    </li>
  </ul>
</el-col>
<el-col :span="8">
</el-col>
</el-row>
</el-main>
</template>

<script>
import axios from '../helpers/axios'

export default {
  data() {
    return {
      todoList: [],
      title: '',
      description: '',
      selectedTodo: null
    }
  },
  methods: {
    //Adds new todo
    addTodo: function() {
      if(this.$data.selectedTodo===null) {
        axios.post('http://127.0.0.1:8000/todo/add', {
          title: this.$data.title,
          description: this.$data.description,
          done: false
        }).then(function(response) {
        //this.$data.todoList.push(response.data);
        console.log(response);        
        });
      } else {
        //Updates an existing todo
         var self = this;
        //var index = this.$data.todoList.indexOf(this.$data.selectedTodo);
        axios.put('http://127.0.0.1:8000/todo/'+ this.$data.selectedTodo.id, {
          id: this.$data.selectedTodo.id,
          title: this.$data.title,
          description: this.$data.description          
        }).then(function(response) {

          console.log(response);

          self.$data.todoList.splice(self.$data.todoList.indexOf(self.$data.selectedTodo), 1);
          self.$data.selectedTodo = null;
          self.$data.title = '';
          self.$data.description = '';
          // console.log(self.$data.todoList);

        });     
      }
    },

    // addTodo: function() {
    //   if(this.$data.selectedTodo===null) {
    //     this.$data.todoList.push({
    //       // id: this.$data.todo.length+1,
    //       title: this.$data.title,
    //       description: this.$data.description
    //     });        
    //   } else {
    //     //var index = this.$data.todoList.indexOf(this.$data.selectedTodo);
    //     this.$data.todoList.splice(this.$data.todoList.indexOf(this.$data.selectedTodo), 1, {
    //       id: this.$data.selectedTodo.id,
    //       title: this.$data.title,
    //       description: this.$data.description          
    //     });
    //     this.$data.selectedTodo = null;
    //     this.$data.title = '';
    //     this.$data.description = '';
    //     console.log(this.$data.todoList);
    //   }
    // },

    //Deleted a todo
    deleteTodo: function(todo) {
      var self = this;
      axios.delete('http://127.0.0.1:8000/todo/'+todo.id+'/delete').then(function(response){
        self.$data.todoList.splice(self.$data.todoList.indexOf(todo), 1);
      }) 
      .catch(function (error) {
self.$data.todoList.splice(self.$data.todoList.indexOf(todo), 1);    console.log(error);
  });
    },
    //Lists all todos
    showTodo: function(todo) {
      this.$data.title = todo.title;
      this.$data.description = todo.description;
      this.$data.selectedTodo = todo;
    }
  },
  mounted: function() {
    var self = this;

    axios.get('http://127.0.0.1:8000/todo/all')
      .then(function (response) {
        // handle success
        self.$data.todoList = response.data;
        console.log(response);
      });
  }
}

</script>

