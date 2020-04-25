https://github.com/golang/go/issues/38617

runtime: make the scavenger's pacing logic more defensive

This change adds two bits of logic to the scavenger's pacing. Firstly,
it checks to make sure we scavenged at least one physical page, if we
released a non-zero amount of memory. If we try to release less than one
physical page, most systems will release the whole page, which could
lead to memory corruption down the road, and this is a signal we're in
this situation.

Secondly, the scavenger's pacing logic now checks to see if the time a
scavenging operation takes is measured to be exactly zero or negative.
The exact zero case can happen if time update granularity is too large
to effectively capture the time the scavenging operation took, like on
Windows where we set the OS timer interrupt frequency to 1 ms. The
negative case should not happen, but we're being defensive (against
kernel bugs, bugs in the runtime, etc.). If either of these cases
happen, we fall back to Go 1.13 behavior: assume the scavenge
operation took around 10µs per physical page. We ignore huge pages in
this case because we're in unknown territory, so we choose to be
conservative about pacing (huge pages could only increase the rate of
scavenging).

Currently, the scavenger is broken on Windows because the granularity of
time measurement is around 1 ms, which is too coarse to measure how fast
we're scavenging, so we often end up with a scavenging time of zero,
followed by NaNs and garbage values in the pacing logic, which usually
leads to sleeping forever.

Fixes [#38617](https://github.com/golang/go/issues/38617).