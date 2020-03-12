# Java vs. C#

## 多线程
[对比Java和.NET多线程编程](https://mp.weixin.qq.com/s/HesvRgE9IozUmh1xDGathw)

![](../img/java-vs-net.webp)

-|Java | .NET
---|---|---
开发语言|java(Scalar etc.)|C#(VB,F# etc.)



## Code for Code Comparison
http://www.javacamp.org/javavscsharp
<table border="1" align="center">
<tbody><tr><td><a href="getStarted.html">Getting started</a></td>
<td><a href="keyword.html">Keywords</a></td>
<td><a href="types.html">Types</a></td>
<td><a href="access.html">Access Modifiers</a></td>
<td><a href="startup.html">Application Startup</a></td>
</tr><tr>
<td><a href="arraytype.html">Array type</a></td>
<td><a href="condition.html">Conditional Statements</a></td>
<td><a href="constructor.html">Constructor</a></td>
<td><a href="default.html">default Modifier</a></td>
<td><a href="documentation.html">Documentation</a></td>
</tr><tr>
<td><a href="eventhandler.html">Event Handler</a></td>
<td><a href="exception.html">Exception</a></td>
<td><a href="sealed.html">final vs. sealed</a></td>
<td><a href="string.html">String vs. string</a></td>
<td><a href="inheritance.html">Inheritance</a></td>
</tr><tr>
<td><a href="instance.html">Instance</a></td>
<td><a href="interface.html">Interface</a></td>
<td><a href="internal.html">internal Modifier</a></td>
<td><a href="keyword.html">Keywords</a></td>
<td><a href="library.html">Library &amp; Package</a></td>
</tr><tr>
<td><a href="loop.html">Loop Constructs</a></td>
<td><a href="method.html">Method Modifiers</a></td>
<td><a href="operators.html">Operators</a></td>
<td><a href="outparam.html">Out Parameters</a></td>
<td><a href="override.html">Overriding</a></td>
</tr><tr>
<td><a href="paramarray.html">Parameter Array</a></td>
<td><a href="prim2obj.html">Primitives vs Objects</a></td>
<td><a href="protected.html">protected Modifier</a></td>
<td><a href="refparam.html">Reference Parameters</a></td>
<td><a href="scope.html">Scope</a></td>
</tr><tr>
<td><a href="stack.html">Stack class</a></td>
<td><a href="static.html">Static Context</a></td>
<td><a href="struct.html">struct</a></td>
<td><a href="sync.html">synchronized vs lock</a></td>
<td><a href="trycatch.html">try/catch Statement</a></td>
</tr><tr>

<td><a href="typeexample.html">Type Example</a></td>
<td><a href="valuetypes.html">Value Type</a></td>
<td><a href="thread.html">Multiple Thread</a></td>
<td><a href="attribute.html">Attribute</a></td>
<td><a href="share.html">Shared Features</a></td>
</tr><tr>

<td><a href="readonly.html">readonly keyword</a></td>
<td><a href="protected.html">protected Modifier</a></td>
<td><a href="property.html">Properties</a></td>
<td><a href="poly.html">Polymorphism</a></td>
<td><a href="pointer.html">Pointer</a></td>
</tr><tr>

<td><a href="passbyvalue.html">Pass By Value vs. Pass By Reference</a></td>
<td><a href="ooloading.html">Operator Overloading</a></td>
<td><a href="objects.html">Objects</a></td>
<td><a href="new.html">new keyword</a></td>
<td><a href="namespace.html">namespace vs. package</a></td>
</tr><tr>

<td><a href="maxmin.html">Maximums and Minmums</a></td>
<td><a href="boxing.html">Boxing vs. Unboxing</a></td>
<td><a href="library.html">Library &amp; Package</a></td>
<td><a href="iniparent.html">Initialization of Parent Constructor</a></td>
<td><a href="indexer.html">Indexer</a></td>
</tr><tr>

<td><a href="foreach.html">foreach</a></td>
<td><a href="fieldmodifier.html">Field modifiers</a></td>
<td><a href="escapesqnce.html">Escape Sequences</a></td>
<td><a href="enum.html">enum</a></td>
<td><a href="event.html">event</a></td>
</tr><tr>

<td><a href="delegate.html">delegate</a></td>
<td><a href="string2num.html">Converting a string to a number</a></td>
<td><a href="readonly.html">const keyword</a></td>
<td><a href="comminter.html">Command-line Interactive</a></td>
<td><a href="classmodifiers.html">Class Modifiers</a></td>
</tr><tr>

<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
</tr><tr>


</tr>
</tbody></table>


## Keywords & Reserved Words
http://www.javacamp.org/javavscsharp/keyword.html

<table align="center"><tbody><tr><td valign="top">
<p align="center"><table class="java" border="1">
<tbody><tr><th><font size="+1">Java</font></th></tr></tbody><tbody><tr><td width="100%"><pre> 
52
abstract     do         if           package      synchronized 
boolean      double     implements   private      this 
break        else       import       protected    throw 
byte         extends    instanceof   public       throws 
case         false      int          return       transient 
catch        final      interface    short        true 
char         finally    long         static       try 
class        float      native       strictfp     void 
const        for        new          super        volatile 
continue     goto       null         switch       while 
default      assert(1.4)             enum(1.5)





<hr>
C# doesn't have
assert       boolean    extends      final        implements
instanceof   native     package      strictfp     super
synchronized            throws       transient





</pre></td></tr></tbody></table>
</p></td><td valign="top">
<p align="center"><table class="csharp" border="1">
<tbody><tr><th><font size="+1">C#</font></th></tr></tbody><tbody><tr><td width="100%"><pre> 
77
abstract     as         base         bool         break
byte         case       catch        char         checked
class        const      continue     decimal      default
delegate     do         double       else         enum
event        explicit   extern       false        finally
fixed        float      for          foreach      goto
if           implicit   in           int          interface
internal     is         lock         long         namespace
new          null       object       operator     out
override     params     private      protected    public
readonly     ref        return       sbyte        sealed
short        sizeof     stackalloc   static       string
struct       switch     this         throw        true
try          typeof     uint         ulong        unchecked
unsafe       ushort     using        virtual      void
volatile     while
<hr>
Java doesn't have
as           base       bool         checked      decimal
delegate                event        explicit     extern
fixed        foreach    implicit     int          internal
is           lock       namespace    object       operator
out          override   params       readonly     ref
sbyte        sealed     sizeof       stackalloc   string
struct       typeof     uint         ulong        unchecked
unsafe       ushort     using        virtual
</pre></td></tr></tbody></table>
</p></td></tr>
</tbody></table>