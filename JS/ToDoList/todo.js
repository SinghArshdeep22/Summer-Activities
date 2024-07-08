"use strict";

const todos = [];

const form = document.getElementById("add_todo_form");
const textInput = document.getElementById("todo_text");

function insertTodo(todos, description) {
  todos.push({
    description: description,
    done: false
  });
}

function render(ul, todos) {
  ul.innerHTML = "";

  for (let i = 0; i < todos.length; i++) {
    const todo = todos[i];

    const li = itemTpl.cloneNode(true);
    ul.appendChild(li);

    const chk = li.querySelector("input");
    if (todo.done) {
      li.classList.add("done");

      chk.checked = true;
      chk.disabled = true;

      setTimeout(function () {
        li.remove(); 
      }, 480);
    } else {
      chk.onclick = function () {
        todo.done = true;
        render(ul, todos); 
      };
    }

    const span = li.querySelector("span");
    span.textContent = todo.description;
  }
}

const todoList = document.querySelector(".todo-list");
const itemTpl = todoList.children[1];
todoList.innerHTML = "";

form.onsubmit = function (e) {
  e.preventDefault();
  const description = textInput.value;
  insertTodo(todos, description);

  render(todoList, todos);
};
