package cliedit

// Elvish code for default bindings, assuming the editor ns as the global ns.
const defaultBindingsElv = `
insert:binding = (binding-table [
  &Left=  $move-dot-left~
  &Right= $move-dot-right~

  &Ctrl-Left=  $move-dot-left-word~
  &Ctrl-Right= $move-dot-right-word~
  &Alt-Left=   $move-dot-left-word~
  &Alt-Right=  $move-dot-right-word~
  &Alt-b=      $move-dot-left-word~
  &Alt-f=      $move-dot-right-word~

  &Home= $move-dot-sol~
  &End=  $move-dot-eol~

  &Backspace= $kill-rune-left~
  &Delete=    $kill-rune-right~
  &Ctrl-W=    $kill-word-left~
  &Ctrl-U=    $kill-line-left~
  &Ctrl-K=    $kill-line-right~

  &Alt-,=  $lastcmd:start~
  &Ctrl-R= $histlist:start~
  &Ctrl-L= $location:start~
  &Ctrl-N= $navigation:start~
  &Tab=    $completion:start~
  &Up=     $history:start~

  &Ctrl-D=  $commit-eof~
])

listing:binding = (binding-table [
  &Up=        $listing:up~
  &Down=      $listing:down~
  &Tab=       $listing:down-cycle~
  &Shift-Tab= $listing:up-cycle~
  &Ctrl-'['=  $close-listing~
])

navigation:binding = (binding-table [
  &Ctrl-'['= $close-listing~
])

completion:binding = (binding-table [
  &Ctrl-'['= $completion:close~
])

history:binding = (binding-table [
  &Up=       $history:up~
  &Down=     $history:down~
  &Ctrl-'['= $history:close~
])

#  &Up=        $listing:up~
#  &Down=      $listing:down~
#  &Tab=       $listing:down-cycle~
#  &Shift-Tab= $listing:up-cycle~
#
#  &Ctrl-F=    $listing:toggle-filtering~
#
#  &Alt-Enter= $listing:accept~
#  &Enter=     $listing:accept-close~
#  &Alt-,=     $listing:accept-close~
#  &Ctrl-'['=  $reset-mode~
#
#  &Default=   $listing:default~
`

// vi: set et:
