<template>
<el-main>
<el-row type="flex" class="row-bg" justify="center">
<el-col :span="8">
</el-col>
<el-col :span="8">
  <el-form ref="form" :model="form" label-width="120px">
  <el-form-item label="Todo">
    <el-input placeholder="Add todo" v-model="todo"></el-input>
  </el-form-item>
  <el-form-item label="Info">
    <el-input placeholder="Todo info" v-model="info"></el-input>
  </el-form-item>
  <el-form-item>
  <el-button type="primary" @click="addTodo" v-if="selectedTodo!==null">Save</el-button>
  <el-button type="primary" @click="addTodo" v-if="selectedTodo===null">Add</el-button>
  </el-form-item>
  </el-form>
  <ul>
    <li v-for="todo in todoList" @click="showTodo(todo)"><b>{{todo.text}}</b> <small>{{todo.info}}</small>
    <el-button @click="deleteTodo(todo)">X</el-button>
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
      todoList: [
        {
          id: 1,
          text: 'test',
          info: 'text text'
        },
        {

          id: 2,
          text: 'test2',
          info: 'text text text'
        }
      ],
      todo: '',
      info: '',
      selectedTodo: null
    }
  },
  methods: {
    addTodo: function() {
      if(this.$data.selectedTodo===null) {
        this.$data.todoList.push({
          id: this.$data.todo.length+1,
          text: this.$data.todo,
          info: this.$data.info
        });        
      } else {
        //var index = this.$data.todoList.indexOf(this.$data.selectedTodo);
        this.$data.todoList.splice(this.$data.todoList.indexOf(this.$data.selectedTodo), 1, {
          id: this.$data.selectedTodo.id,
          text: this.$data.todo,
          info: this.$data.info          
        });
        this.$data.selectedTodo = null;
        this.$data.todo = '';
        this.$data.info = '';
        console.log(this.$data.todoList);
      }
    },
    deleteTodo: function(todo) {
      this.$data.todoList.splice(this.$data.todoList.indexOf(todo), 1);
    },
    showTodo: function(todo) {
      this.$data.todo = todo.text;
      this.$data.info = todo.info;
      this.$data.selectedTodo = todo;
    }
  },
  // mounted: function() {
  //   var self = this;

  // axios.get('https://jsonplaceholder.typicode.com/todos')
  //   .then(function (response) {
  //     // handle success
  //     self.$data.todoList = response.data;
  //     console.log(response);
  //   });
  // }
}

</script>
