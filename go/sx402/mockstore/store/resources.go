package store

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/coinbase/x402/go/pkg/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/san-lab/sx402/state"
)

// TemplateData defines the data passed to the HTML template
type TemplateData struct {
	TxHash      string
	Explorer    string
	Content     template.HTML
	Network     string
	Status      string
	Facilitator string
}

func ResourceHandler(c *gin.Context) {
	// Parse RESID
	resid := c.Query("RESID")
	idx, err := strconv.Atoi(resid)
	if err != nil || idx < 1 || idx > len(Stories) {
		c.String(http.StatusNotFound, "Invalid resource ID")
		return
	}

	// Retrieve tx_hash from context (middleware must set this beforehand)
	settleResp0, exists := c.Get("settleReponse")
	txHash := "(unknown)" // fallback if not set
	if exists {
		settleResponse := settleResp0.(*types.SettleResponse)
		txHash = settleResponse.Transaction
	}

	// Parse template
	tmpl, err := template.ParseFiles(StorePrefix + "/templates/story.tmpl")
	if err != nil {
		c.String(http.StatusInternalServerError, "Template error: %v", err)
		return
	}

	explorer := c.GetString("explorer")

	// Render template
	data := TemplateData{
		TxHash:      txHash,
		Explorer:    explorer,
		Content:     template.HTML(Stories[idx-1]),
		Network:     c.GetString("network"),
		Facilitator: "/facilitator/receipt",
		Status:      "Unknown",
	}

	time.Sleep(time.Second)
	_, status := state.GetPendingReceipt(common.HexToHash(txHash), c.GetString("network"))

	if status {
		data.Status = "Settled"
	}

	c.Status(http.StatusOK)
	tmpl.Execute(c.Writer, data)
}

var Stories = []string{`<div class="story">
    <h2>The King</h2>
    <p>
      Once, in a peaceful kingdom nestled between the mountains and the sea, there lived a king who spoke little but thought deeply. He wore no crown, sat on no throne, and walked among his people in simple robes.
    </p>
    <p>
      One morning, an advisor asked, “Your Majesty, why do you not rule as other kings do—with commands and glory?”
    </p>
    <p>
      The king smiled. “Because I rule best when I listen most.”
    </p>
    <p>
      He spent his days in the markets, at the docks, and in the fields. He listened to bakers, sailors, and shepherds. When he finally made decisions, they were wise—because they weren’t just his own.
    </p>
    <p>
      The kingdom prospered not through fear or force, but because the people felt heard. And so the king, quiet and thoughtful, was remembered for generations—not for the wars he fought, but for the peace he kept.
    </p>
  </div>`,
	`<div class="story">
    <h2>The Princess</h2>
    <p>
      The princess had a secret: she could talk to birds. Not just chirps and songs, but real conversations—about wind, about trees, about everything they saw from the sky.
    </p>
    <p>
      Each morning, she’d sneak into the royal garden and sit among the vines. The sparrows would tell her where ships were coming from, the swans described the moods of the lake, and the crows gossiped about palace guards.
    </p>
    <p>
      One day, a foreign prince came to ask for her hand. He arrived with jewels and speeches. The princess said nothing—only listened. That evening, a raven landed on her shoulder and whispered, “He lies.”
    </p>
    <p>
      She declined the proposal politely but firmly. The court was shocked. “How did you know?” they asked.
    </p>
    <p>
      She smiled and said, “The sky speaks, if you learn to listen.”
    </p>
    <p>
      And so the princess ruled one day not just with grace, but with a wisdom that came from above.
    </p>
  </div>`,
	`<div class="story">
    <h2>The Jester</h2>
    <p>
      The court jester was a fool by title but a philosopher by heart. With painted face and jingling cap, he danced, joked, and played the lute—but always with a wink of something deeper.
    </p>
    <p>
      When the king was troubled, it was the jester who made him laugh. When the queen wept, it was the jester who sang a song about stars falling in love with rivers.
    </p>
    <p>
      One day, a nobleman mocked him: “You are nothing but a clown.”
    </p>
    <p>
      The jester bowed. “True! But even a clown can hold a mirror.”
    </p>
    <p>
      That night, his play was about a kingdom that lost its way because it stopped laughing. The audience chuckled at first. Then they thought. Then they clapped.
    </p>
    <p>
      The jester left the stage with a smile. He knew that the hardest truths were best hidden inside a joke—just sharp enough to tickle the soul.
    </p>
  </div>`}
