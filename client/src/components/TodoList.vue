<template>
    <div>
        <header class="bg-primary text-center text-light p-3">
            <h1>
                Todo
                <span class="info">({{ remaining.length }}/{{ items.length }})</span>
            </h1>
        </header>
        <div class="container mt-5">
            <div class="input-group mt-5 mb-4">
                <label class="input-group-text">タスク名</label>
                <input type="text" v-model="content" class="form-control">
                <label class="input-group-text">優先度</label>
                <select class="form-select" v-model="priority">
                    <option selected>優先度</option>
                    <option v-for="(badge, index) in badges" :key="index" :value="index">{{badge.name}}</option>
                </select>
                <button @click="addItem" type="button" class="btn btn-primary">追加</button>
            </div>
            <div class="list-group">
                <div v-for="(todo, index) in items" :key="index" :class="{'bg-secondary': todo.is_done}" :value="index" class="list-group-item list-group-item-action">
                    <TodoListItem 
                        v-if="index !== edit_index"
                        v-bind="todo"
                        v-on:done-item="doneItem"
                        v-on:delete-item="deleteItem"
                        v-on:edit-item="editItem"/>
                    <TodoListEdit
                        v-else
                        v-bind="todo"
                        v-on:save-item="saveItem" />
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import api from '../services/api.js'
import TodoListItem from './TodoListItem.vue'
import TodoListEdit from './TodoListEdit.vue'
import { Badges } from '../services/badges.js'

export default {
    name: 'TodoList',
    data: function() {
        return {
            content: '',
            priority: '',
            items: [],
            edit_index: -1,
            badges: Badges
        }
    },
    components: {
        TodoListItem,
        TodoListEdit
    },
    mounted: async function() {
        this.items = await api.getlist();
    },
    methods: {
        addItem: async function() {
            this.items = await api.regist(this.content, parseInt(this.priority));
            this.content = '';
            this.priority = '';
        },
        deleteItem: async function(id) {
            if(confirm('削除しますか?')) {
                let res = await api.delete(id);
                this.items = this.items.filter(todo => todo.id != res.id);
            }
        },
        doneItem: async function(id, done) {
            var index = this.items.findIndex(t => t.id === id);
            this.items[index].is_done = done;
            api.update(this.items[index]);
        },
        editItem: function(id) {
            this.edit_index = this.items.findIndex(t => t.id === id);
        },
        saveItem: function(id, content, priority) {
            var index = this.items.findIndex(t => t.id === id);
            this.items[index].content = content;
            this.items[index].priority = priority;
            api.update(this.items[index]);

            this.edit_index = -1;
        }
    },
    computed: {
        remaining: function() {
            return this.items.filter(function(todo) {
                return !todo.is_done;
            });
        }
    }
}
</script>
