const API_URL = "https://YOUR-ULTRON-URL.onrender.com";

const chat = document.getElementById("chat");
const form = document.getElementById("chat-form");
const input = document.getElementById("message-input");
const sendButton = document.getElementById("send-button");
const emptyState = document.getElementById("empty-state");

const newChatButton = document.getElementById("new-chat");
const clearChatButton = document.getElementById("clear-chat");

let messages = [];


/* =========================
   SEND MESSAGE
========================= */

form.addEventListener("submit", async (event) => {
    event.preventDefault();

    const message = input.value.trim();

    if (!message) {
        return;
    }

    addMessage("user", message);

    input.value = "";
    autoResize();

    setLoading(true);

    const typing = addTypingIndicator();

    try {
        const response = await fetch(`${API_URL}/chat`, {
            method: "POST",

            headers: {
                "Content-Type": "application/json"
            },

            body: JSON.stringify({
                message: message
            })
        });

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`);
        }

        const data = await response.json();

        typing.remove();

        addMessage(
            "assistant",
            data.response || "I didn't receive a response."
        );

    } catch (error) {
        typing.remove();

        console.error(error);

        addMessage(
            "assistant",
            "I couldn't connect to Ultron right now. Please try again."
        );
    }

    setLoading(false);
});


/* =========================
   ADD MESSAGE
========================= */

function addMessage(role, text) {

    emptyState.style.display = "none";

    const messageElement = document.createElement("div");

    messageElement.className = `message ${role}`;

    const content = document.createElement("div");

    content.className = "message-content";

    content.textContent = text;

    messageElement.appendChild(content);

    chat.appendChild(messageElement);

    messages.push({
        role,
        text
    });

    scrollToBottom();
}


/* =========================
   TYPING INDICATOR
========================= */

function addTypingIndicator() {

    emptyState.style.display = "none";

    const message = document.createElement("div");

    message.className = "message assistant";

    message.innerHTML = `
        <div class="message-content">
            <div class="typing">
                <span></span>
                <span></span>
                <span></span>
            </div>
        </div>
    `;

    chat.appendChild(message);

    scrollToBottom();

    return message;
}


/* =========================
   LOADING
========================= */

function setLoading(loading) {

    sendButton.disabled = loading;

    input.disabled = loading;

    if (!loading) {
        input.focus();
    }
}


/* =========================
   AUTO RESIZE
========================= */

input.addEventListener("input", autoResize);

function autoResize() {

    input.style.height = "auto";

    input.style.height =
        Math.min(input.scrollHeight, 180) + "px";
}


/* =========================
   ENTER TO SEND
========================= */

input.addEventListener("keydown", (event) => {

    if (event.key === "Enter" && !event.shiftKey) {

        event.preventDefault();

        form.requestSubmit();
    }
});


/* =========================
   SCROLL
========================= */

function scrollToBottom() {

    requestAnimationFrame(() => {

        chat.scrollTo({
            top: chat.scrollHeight,
            behavior: "smooth"
        });

    });
}


/* =========================
   NEW CHAT
========================= */

newChatButton.addEventListener("click", clearConversation);

clearChatButton.addEventListener("click", clearConversation);

function clearConversation() {

    messages = [];

    chat.innerHTML = "";

    chat.appendChild(emptyState);

    emptyState.style.display = "";

    input.value = "";

    autoResize();

    input.focus();
}


/* =========================
   SUGGESTIONS
========================= */

document.querySelectorAll(".suggestion").forEach((button) => {

    button.addEventListener("click", () => {

        const text = button.textContent.trim();

        if (text.includes("Explain")) {
            input.value = "Explain something interesting to me.";
        }

        if (text.includes("code")) {
            input.value = "Help me write some code.";
        }

        if (text.includes("Brainstorm")) {
            input.value = "Help me brainstorm a project idea.";
        }

        autoResize();

        input.focus();
    });

});
