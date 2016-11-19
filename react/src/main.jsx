var Comment = React.createClass({
    render: function() {
        return (
        <div className="comment">
            <h2 className="commentAuthor">
                作者:{this.props.author}
            </h2>
               内容:{this.props.children}
        </div>
        );
    }
    });
  var CommentForm = React.createClass({
    getInitialState: function() {
    return {author: '', text: ''};
    },
    handleAuthorChange: function(e) {
        this.setState({author: e.target.value});
    },
    handleTextChange: function(e) {
        this.setState({text: e.target.value});
    },
    
     handleSubmit: function(e) {
        e.preventDefault();
        var author = this.state.author.trim();
        var text = this.state.text.trim();
        if (!text || !author) {
        return;
        }
        // TODO: send request to the server
        console.log('Ìá½»µ½·þÎñÆ÷');
        this.props.onCommentSubmit({author: author, text: text});
        this.setState({author: '', text: ''});
        
    },
    render: function() {
        return (
        <form className="commentForm"  onSubmit={this.handleSubmit}>
            <input type="text" placeholder="Your name" 
            value={this.state.author}
            onChange={this.handleAuthorChange}
            />
            <br />
            <input type="text" placeholder="Say something..."
            value={this.state.text}
            onChange={this.handleTextChange} />
            <input type="submit" value="Post" />
        </form>
        );
    }
    });
    var data = [
        {id: 1, author: "Pete Hunt", text: "This is one comment"},
        {id: 2, author: "Jordan Walke", text: "This is *another* comment"}
    ];
    var CommentItem = React.createClass({
        render: function() {
            console.log(this.props.data);
            var commentNodes = this.props.data.map(function(comment) {
                return (
                    <Comment author={comment.author} key={comment.id}>
                        {comment.text}
                    </Comment>
                );
            });
            return (
            <div className="commentItem">
                {commentNodes}
            </div>
            );
        }
    });
    var CommentBox = React.createClass({
    	
         getInitialState: function() {
        
            return {
                data: []
            };
        },
        handleCommentSubmit: function(comment) {
               
                var comments = this.state.data;
                comment.id = Date.now();
                var newComments = comments.concat([comment]);
                this.setState({data: newComments});
        },
        componentDidMount:function(){
           
        },
        render: function() {
          
            console.log(this.state.data);
            return (
            <div className="commentBox">
                <CommentItem data={this.state.data}  />
                <CommentForm  onCommentSubmit={this.handleCommentSubmit} />
            </div>
            );
        }
    });
    ReactDOM.render(
    <CommentBox />,
    document.getElementById('content')
    );