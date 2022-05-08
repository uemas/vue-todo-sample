import axios from "axios";

const client = axios.create({
    baseURL: `http://localhost:1323/`,
    headers: {
        'Content-Type': 'application/json',
    },
});

export default {
    regist: async function(content, priority) {
        return client.post('regist', JSON.stringify({
            "content": content,
            "priority": priority,
            "is_done": false
        })).then(res => res.data);
    },
    delete: async function(id) {
        return client.delete(`delete/${id}`).then(res => res.data);
    },
    update: async function(todo) {
        return client.put(`update/${todo.id}`, JSON.stringify({
            "content": todo.content,
            "priority": todo.priority,
            "is_done": todo.is_done
        })).then(res => res.data);
    },
    getlist: async function() {
        return client.get('list').then(res => res.data);
    }
}
