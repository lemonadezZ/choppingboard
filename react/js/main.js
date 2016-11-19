var Comment = React.createClass({
    displayName: "Comment",

    render: function () {
        return React.createElement(
            "div",
            { className: "comment" },
            React.createElement(
                "h2",
                { className: "commentAuthor" },
                "\u4F5C\u8005:",
                this.props.author
            ),
            "\u5185\u5BB9:",
            this.props.children
        );
    }
});
var CommentForm = React.createClass({
    displayName: "CommentForm",

    getInitialState: function () {
        return { author: '', text: '' };
    },
    handleAuthorChange: function (e) {
        this.setState({ author: e.target.value });
    },
    handleTextChange: function (e) {
        this.setState({ text: e.target.value });
    },

    handleSubmit: function (e) {
        e.preventDefault();
        var author = this.state.author.trim();
        var text = this.state.text.trim();
        if (!text || !author) {
            return;
        }
        // TODO: send request to the server
        console.log('Ìá½»µ½·þÎñÆ÷');
        this.props.onCommentSubmit({ author: author, text: text });
        this.setState({ author: '', text: '' });
    },
    render: function () {
        return React.createElement(
            "form",
            { className: "commentForm", onSubmit: this.handleSubmit },
            React.createElement("input", { type: "text", placeholder: "Your name",
                value: this.state.author,
                onChange: this.handleAuthorChange
            }),
            React.createElement("br", null),
            React.createElement("input", { type: "text", placeholder: "Say something...",
                value: this.state.text,
                onChange: this.handleTextChange }),
            React.createElement("input", { type: "submit", value: "Post" })
        );
    }
});
var data = [{ id: 1, author: "Pete Hunt", text: "This is one comment" }, { id: 2, author: "Jordan Walke", text: "This is *another* comment" }];
var CommentItem = React.createClass({
    displayName: "CommentItem",

    render: function () {
        console.log(this.props.data);
        var commentNodes = this.props.data.map(function (comment) {
            return React.createElement(
                Comment,
                { author: comment.author, key: comment.id },
                comment.text
            );
        });
        return React.createElement(
            "div",
            { className: "commentItem" },
            commentNodes
        );
    }
});
var CommentBox = React.createClass({
    displayName: "CommentBox",


    getInitialState: function () {

        return {
            data: []
        };
    },
    handleCommentSubmit: function (comment) {

        var comments = this.state.data;
        comment.id = Date.now();
        var newComments = comments.concat([comment]);
        this.setState({ data: newComments });
    },
    componentDidMount: function () {},
    render: function () {

        console.log(this.state.data);
        return React.createElement(
            "div",
            { className: "commentBox" },
            React.createElement(CommentItem, { data: this.state.data }),
            React.createElement(CommentForm, { onCommentSubmit: this.handleCommentSubmit })
        );
    }
});
ReactDOM.render(React.createElement(CommentBox, null), document.getElementById('content'));