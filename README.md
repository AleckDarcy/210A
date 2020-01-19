# install antlr4
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.7.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java -Xmx500M -cp "/usr/local/lib/antlr-4.7.2-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'

# generate code for grammar
antlr4 -Dlanguage=Go -package grammar zz/grammar/ZZ.g4
