// Example JavaScript Application: Task Manager

// Data Model
class Task {
  constructor(id, title, description = '', dueDate = null, completed = false) {
    this.id = id;
    this.title = title;
    this.description = description;
    this.dueDate = dueDate;    // ISO string
    this.completed = completed;
  }
}

// Task Store (in-memory for demo)
class TaskStore {
  constructor() {
    this.tasks = [];
    this.nextId = 1;
  }

  addTask(title, description, dueDate) {
    const task = new Task(this.nextId++, title, description, dueDate);
    this.tasks.push(task);
    return task;
  }

  getTask(id) {
    return this.tasks.find(t => t.id === id) || null;
  }

  updateTask(id, fields) {
    const task = this.getTask(id);
    if (!task) throw new Error(`Task ${id} not found`);
    Object.assign(task, fields);
    return task;
  }

  deleteTask(id) {
    const idx = this.tasks.findIndex(t => t.id === id);
    if (idx === -1) throw new Error(`Task ${id} not found`);
    return this.tasks.splice(idx, 1)[0];
  }

  listTasks(filter = {}) {
    return this.tasks.filter(t => {
      if (filter.completed != null && t.completed !== filter.completed) return false;
      if (filter.dueDate && t.dueDate > filter.dueDate) return false;
      return true;
    });
  }
}

// Utility Functions
function formatDate(isoString) {
  if (!isoString) return '';
  const d = new Date(isoString);
  return d.toLocaleDateString('en-US', {
    year: 'numeric', month: 'short', day: 'numeric'
  });
}

function createElement(tag, attrs = {}, ...children) {
  const el = document.createElement(tag);
  Object.entries(attrs).forEach(([k, v]) => {
    if (k === 'className') el.className = v;
    else if (k.startsWith('on') && typeof v === 'function') {
      el.addEventListener(k.substring(2).toLowerCase(), v);
    } else {
      el.setAttribute(k, v);
    }
  });
  children.forEach(child => {
    if (typeof child === 'string') el.appendChild(document.createTextNode(child));
    else if (child instanceof Node) el.appendChild(child);
  });
  return el;
}

// DOM Rendering
class TaskApp {
  constructor(rootId) {
    this.store = new TaskStore();
    this.root = document.getElementById(rootId);
    this.init();
  }

  init() {
    this.root.innerHTML = '';
    this.form = this.renderForm();
    this.listContainer = createElement('div', { className: 'task-list' });
    this.root.append(this.form, this.listContainer);
    this.bindEvents();
    this.refreshList();
  }

  renderForm() {
    const titleInput = createElement('input', { type: 'text', placeholder: 'Title', id: 'task-title' });
    const descInput  = createElement('textarea', { placeholder: 'Description', id: 'task-desc' });
    const dueInput   = createElement('input', { type: 'date', id: 'task-due' });
    const addBtn     = createElement('button', { type: 'button', onClick: () => this.handleAdd() }, 'Add Task');

    return createElement('div', { className: 'task-form' },
      titleInput, descInput, dueInput, addBtn
    );
  }

  bindEvents() {
    document.addEventListener('task:added', () => this.refreshList());
    document.addEventListener('task:updated', () => this.refreshList());
    document.addEventListener('task:deleted', () => this.refreshList());
  }

  handleAdd() {
    const title = document.getElementById('task-title').value.trim();
    const desc  = document.getElementById('task-desc').value.trim();
    const due   = document.getElementById('task-due').value || null;
    if (!title) return alert('Title is required');

    this.store.addTask(title, desc, due);
    document.dispatchEvent(new CustomEvent('task:added'));
    this.clearForm();
  }

  handleToggleComplete(id) {
    const task = this.store.getTask(id);
    this.store.updateTask(id, { completed: !task.completed });
    document.dispatchEvent(new CustomEvent('task:updated'));
  }

  handleDelete(id) {
    if (!confirm('Delete this task?')) return;
    this.store.deleteTask(id);
    document.dispatchEvent(new CustomEvent('task:deleted'));
  }

  clearForm() {
    document.getElementById('task-title').value = '';
    document.getElementById('task-desc').value  = '';
    document.getElementById('task-due').value   = '';
  }

  renderTaskItem(task) {
    const title = createElement('span', { className: task.completed ? 'completed' : '' }, task.title);
    const due   = createElement('small', {}, formatDate(task.dueDate));
    const toggleBtn = createElement('button', {
      onClick: () => this.handleToggleComplete(task.id)
    }, task.completed ? 'Undo' : 'Complete');
    const delBtn = createElement('button', {
      onClick: () => this.handleDelete(task.id)
    }, 'Delete');

    return createElement('div', { className: 'task-item' },
      title, due, toggleBtn, delBtn
    );
  }

  refreshList() {
    this.listContainer.innerHTML = '';
    const tasks = this.store.listTasks();
    if (tasks.length === 0) {
      this.listContainer.append(createElement('p', {}, 'No tasks yet.'));
      return;
    }
    tasks.forEach(task => {
      this.listContainer.append(this.renderTaskItem(task));
    });
  }
}

// Initialize the app on DOMContentLoaded
document.addEventListener('DOMContentLoaded', () => {
  new TaskApp('app');
});
