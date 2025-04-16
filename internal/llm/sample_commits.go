package llm

import "math/rand"

var commits []string = []string{
	`fs/pipe: fix pipe buffer index use in FUSE
This was another case that Rasmus pointed out where the direct access to
the pipe head and tail pointers broke on 32-bit configurations due to
the type changes.

As with the pipe FIONREAD case, fix it by using the appropriate helper
functions that deal with the right pipe index sizing.
Reported-by: Rasmus Villemoes <ravi@prevas.dk>
Link: https://lore.kernel.org/all/878qpi5wz4.fsf@prevas.dk/
Fixes: 3d25216 ("fs/pipe: Read pipe->{head,tail} atomically outside pipe->mutex")Cc: Oleg >`,
	`fix: prevent racing of requests

Introduce a request id and a reference to latest request. Dismiss
incoming responses other than from latest request.

Remove timeouts which were used to mitigate the racing issue but are
obsolete now.

Reviewed-by: Z
Refs: #123`,
	`When an AfterFunc executes in a synctest bubble, there is a series of
happens-before relationships:

  1. The AfterFunc is created.
  2. The AfterFunc goroutine executes.
  3. The AfterFunc goroutine returns.
  4. A subsequent synctest.Wait call returns.

We were failing to correctly establish the happens-before relationship
between the AfterFunc goroutine and the AfterFunc itself being created.
When an AfterFunc executes, the G running the timer temporarily switches
to the timer heap's racectx. It then calls time.goFunc, which starts a
new goroutine to execute the timer. time.goFunc relies on the new goroutine
inheriting the racectx of the G running the timer.

Normal, non-synctest timers, execute with m.curg == nil, which causes
new goroutines to inherit the g0 racectx. We were running synctest
timers with m.curg set (to the G executing synctest.Run), so the new
AfterFunc goroutine was created using m.curg's racectx. This resulted
in us not properly establishing the happens-before relationship between
AfterFunc being called and the AfterFunc goroutine starting.

Fix this by setting m.curg to nil while executing timers.

As one additional fix, when waking a blocked bubble, wake the root
goroutine rather than a goroutine blocked in Wait if there is a
timer that can fire.

Fixes #72750`,
	`crypto/tls: relax native FIPS 140-3 mode
We are going to stick to BoringSSL's policy for Go+BoringCrypto, but
when using the native FIPS 140-3 module we can allow Ed25519, ML-KEM,
and P-521.

NIST SP 800-52r2 is stricter, but it only applies to some entities, so
they can restrict the profile with Config.

Fixes #71757`,
	`Make Powershell completion script work in constrained mode (#2196)
Creating CompletionResult objects is not allowed in Powershell constrained mode, so return results as strings if constrained mode is enabled

Store results as PsCustomObjects instead of hashtables. This prevents Sort-Object from trying to convert the hashtable to a object, which is blocked in constrained mode.
PsCustomObjects are created using New-Object to work around PowerShell/PowerShell#20767`,
	`Fix deprecation comment for Command.SetOutput (#2172)
Deprecation comments should be at the start of a paragraph [1], and because
of that have a whitespace above them [2];

> To signal that an identifier should not be used, add a paragraph to its
> doc comment that begins with Deprecated: followed by some information
> about the deprecation (...)

With the whitespace missing, some tools, including pkg.go.dev [3] don't
detect it to be deprecated.

[1]: https://go.dev/wiki/Deprecated
[2]: https://go.dev/doc/comment#paragraphs
[3]: https://pkg.go.dev/github.com/spf13/cobra@v1.8.1#Command.SetOutput`,
	`Fix help text for runnable plugin command
When creating a plugin without sub commands, the help text included the
command name (kubectl-plugin) instead of the display name (kubectl plugin):

    Usage:
      kubectl-plugin [flags]

The issue is that the usage line for this case does not use the
command path but the raw "Use" string, and this case was not tested.

Add a test for this case and fix UsageLine() to replace the command name
with the display name.

Tested using https://github.com/nirs/kubernetes/tree/sample-cli-plugin-help`,
}

func getCimmitExample() string {
	return commits[rand.Intn(len(commits))]
}
