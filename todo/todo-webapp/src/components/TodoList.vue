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
  <ol>
    <li v-for="todo in todoList"><b @click="showTodo(todo)">{{todo.title}}</b> <small>{{todo.description}}</small>
    <el-button @click="deleteTodo(todo)" class="btn">X</el-button>
    </li>
  </ol>
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
         var self = this;
      if(this.$data.selectedTodo===null) {
        axios.post('http://localhost:8000/todo/add', {
          title: this.$data.title,
          description: this.$data.description,
          done: false
        }).then(function(response) {
        self.$data.todoList.push(response.data);
        console.log(response);        
        });
      } else {

        //Updates an existing todo
        axios.put('http://localhost:8000/todo/'+ this.$data.selectedTodo.id, {
          id: this.$data.selectedTodo.id,
          title: this.$data.title,
          description: this.$data.description          
        }).then(function(response) {

          console.log(response);

          self.$data.todoList.splice(self.$data.todoList.indexOf(self.$data.selectedTodo), 1, response.data);
          
          self.$data.selectedTodo = null;
          self.$data.title = '';
          self.$data.description = '';
        });     
      }
    },

    //Delete a todo
    deleteTodo: function(todo) {
      var self = this;
      axios.delete('http://localhost:8000/todo/'+todo.id+'/delete').then(function(response){
        self.$data.todoList.splice(self.$data.todoList.indexOf(todo), 1);
      }) 
      .catch(function (error) {
      self.$data.todoList.splice(self.$data.todoList.indexOf(todo), 1);    
      console.log(error);
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

    axios.get('http://localhost:8000/todo/all')
      .then(function (response) {
        // handle success
        self.$data.todoList = response.data;
        console.log(response);
      });
  }
}

</script>

