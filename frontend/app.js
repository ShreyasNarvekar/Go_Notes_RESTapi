function createDeleteButton(onDelete) {
    const button = document.createElement("button");
    button.type = "button";
    button.className = "delete-btn";
    button.setAttribute("aria-label", "Delete item");
    button.textContent = "Delete";
    button.addEventListener("click", onDelete);
    return button;
}

async function fetchJson(path, options) {
    const response = await fetch(path, options);
    if (!response.ok) {
        const errorText = await response.text();
        throw new Error(errorText || "request failed");
    }
    if (response.status === 204) {
        return null;
    }
    return response.json();
}

async function initNotesPage() {
    const notesList = document.getElementById("notes-list");
    const addNoteBtn = document.getElementById("add-note-btn");
    const noteTitleInput = document.getElementById("new-note-title");
    const noteContentInput = document.getElementById("new-note-content");

    if (!notesList || !addNoteBtn || !noteTitleInput || !noteContentInput) {
        return;
    }

    async function renderNotes() {
        const notes = await fetchJson("/notes");
        notesList.innerHTML = "";

        notes.forEach((note) => {
            const item = document.createElement("li");
            item.className = "item";

            const text = document.createElement("div");
            text.className = "item-text";
            text.textContent = `${note.title || "(untitled)"}: ${note.content || ""}`;

            const actions = document.createElement("div");
            actions.className = "item-actions";

            const deleteBtn = createDeleteButton(async () => {
                const confirmed = window.confirm("Delete this note?");
                if (!confirmed) {
                    return;
                }
                await fetchJson(`/notes/${note.id}`, { method: "DELETE" });
                await renderNotes();
            });

            actions.appendChild(deleteBtn);
            item.appendChild(text);
            item.appendChild(actions);
            notesList.appendChild(item);
        });
    }

    addNoteBtn.addEventListener("click", async () => {
        const title = noteTitleInput.value.trim();
        const content = noteContentInput.value.trim();

        if (!title && !content) {
            return;
        }

        await fetchJson("/notes", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title, content })
        });

        noteTitleInput.value = "";
        noteContentInput.value = "";
        await renderNotes();
    });

    await renderNotes();
}

async function initTasksPage() {
    const tasksList = document.getElementById("tasks-list");
    const addTaskBtn = document.getElementById("add-task-btn");
    const taskTitleInput = document.getElementById("new-task-title");
    const taskDescriptionInput = document.getElementById("new-task-description");

    if (!tasksList || !addTaskBtn || !taskTitleInput || !taskDescriptionInput) {
        return;
    }

    async function renderTasks() {
        const tasks = await fetchJson("/tasks");
        tasksList.innerHTML = "";

        tasks.forEach((task) => {
            const item = document.createElement("li");
            item.className = "item";

            const text = document.createElement("div");
            text.className = "item-text";
            text.textContent = `${task.title || "(untitled)"}: ${task.description || ""}`;

            const actions = document.createElement("div");
            actions.className = "item-actions";

            const deleteBtn = createDeleteButton(async () => {
                const confirmed = window.confirm("Delete this task?");
                if (!confirmed) {
                    return;
                }
                await fetchJson(`/tasks/${task.id}`, { method: "DELETE" });
                await renderTasks();
            });

            actions.appendChild(deleteBtn);
            item.appendChild(text);
            item.appendChild(actions);
            tasksList.appendChild(item);
        });
    }

    addTaskBtn.addEventListener("click", async () => {
        const title = taskTitleInput.value.trim();
        const description = taskDescriptionInput.value.trim();

        if (!title && !description) {
            return;
        }

        await fetchJson("/tasks", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ title, description })
        });

        taskTitleInput.value = "";
        taskDescriptionInput.value = "";
        await renderTasks();
    });

    await renderTasks();
}

window.addEventListener("DOMContentLoaded", async () => {
    try {
        await initNotesPage();
        await initTasksPage();
    } catch (error) {
        console.error(error);
        alert("Unable to load data. Ensure backend is running on the same origin.");
    }
});
