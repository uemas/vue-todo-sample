<template>
    <div class="input-group">
        <label class="input-group-text">タスク名</label>
        <input type="text" v-model="edit_content" class="form-control">
        <label class="input-group-text">優先度</label>
        <select class="form-select" v-model="edit_priority">
            <option selected>優先度</option>
            <option v-for="(badge, index) in badges" :key="index" :value="index">{{badge.name}}</option>
        </select>
        <button @click="saveItem" type="button" class="btn btn-primary">保存</button>
    </div>
</template>

<script>
import { Badges } from '../services/badges.js'

export default {
    name: 'TodoListEdit',
    props: ['id', 'content', 'priority', 'is_done'],
    data: function() {
        return {
            edit_content: this.content,
            edit_priority: this.priority,
            badges: Badges
        }
    },
    emits: ['saveItem'],
    methods: {
        saveItem: function() {
            return this.$emit('saveItem', this.id, this.edit_content, this.edit_priority);
        }
    }
}
</script>
