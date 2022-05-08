<template>
    <label class="d-flex justify-content-between align-items-center">
        <div>
            <input type="checkbox" v-model="done" class="form-check-input me-2">
            <span class="me-4" :class="{'text-decoration-line-through': done}">{{content}}</span>
            <span :class="[badge_bg]" class="badge rounded-pill">{{badge_name}}</span>
        </div>
        <div class="btn-group">
            <button @click="editItem(id)" class="btn btn-success">変更</button>
            <button @click="deleteItem(id)" class="btn btn-danger">削除</button>
        </div>
    </label>
</template>

<script>
import { Badges } from '../services/badges.js'

export default {
    name: 'TodoListItem',
    props: ['id', 'content', 'priority', 'is_done'],
    data: function() {
        return {
            done: this.is_done,
            badge_bg: '',
            badge_name: '',
        }
    },
    emits: ['doneItem','deleteItem', 'editItem'],
    watch: {
        done: {
            handler() {
                this.$emit('doneItem', this.id, this.done);
            }
        }
    },
    mounted: function() {
        this.registBadge();
    },
    methods: {
        registBadge: function() {
            let badge = Badges[this.priority];
            this.badge_name = badge.name;
            this.badge_bg = badge.bg;
        },
        deleteItem: function(id) {
            this.$emit('deleteItem', id);
        },
        editItem: function(id) {
            this.$emit('editItem', id);
        }
    }
}
</script>
